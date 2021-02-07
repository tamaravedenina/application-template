package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/utrack/clay/v2/transport"
)

// Server is a transport server.
type Server struct {
	opts      *serverOpts
	listeners *listenerSet
	srv       *serverSet
}

// NewServer creates a Server listening on the rpcPort.
// Pass additional Options to mutate its behaviour.
// By default, HTTP JSON handler and gRPC are listening on the same
// port, admin port is p+2 and profile port is p+4.
func NewServer(grpcPort int, httpPort int, opts ...Option) *Server {
	serverOpts := defaultServerOpts(grpcPort, httpPort)
	for _, opt := range opts {
		opt(serverOpts)
	}
	return &Server{opts: serverOpts}
}

// Run starts processing requests to the service.
// It blocks indefinitely, run asynchronously to do anything after that.
func (s *Server) Run(svc ...transport.Service) error {
	descSlice := make([]transport.ServiceDesc, 0)
	for _, service := range svc {
		descSlice = append(descSlice, service.GetDescription())
	}

	var err error
	s.listeners, err = newListenerSet(s.opts)
	if err != nil {
		return errors.Wrap(err, "couldn't create listeners")
	}

	s.srv = newServerSet(s.opts)

	for _, desc := range descSlice {
		// apply gRPC interceptor
		if d, ok := desc.(transport.ConfigurableServiceDesc); ok {
			d.Apply(transport.WithUnaryInterceptor(s.opts.GRPCUnaryInterceptor))
		}

		// Register everything
		desc.RegisterHTTP(s.srv.http)
		desc.RegisterGRPC(s.srv.grpc)
	}

	// Inject static Swagger as root handler
	s.srv.http.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		swgBytes, _ := generateSwagger(descSlice)
		io.Copy(w, bytes.NewReader(swgBytes))
	})

	return s.run()
}

func (s *Server) run() error {
	errChan := make(chan error, 5)

	if s.listeners.mainListener != nil {
		go func() {
			err := s.listeners.mainListener.Serve()
			errChan <- err
		}()
	}
	go func() {
		err := http.Serve(s.listeners.HTTP, s.srv.http)
		errChan <- err
	}()
	go func() {
		err := s.srv.grpc.Serve(s.listeners.GRPC)
		errChan <- err
	}()

	return <-errChan
}

// Stop stops the server gracefully.
func (s *Server) Stop() {
	// TODO grace HTTP
	s.srv.grpc.GracefulStop()
}
