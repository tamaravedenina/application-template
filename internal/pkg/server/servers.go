package server

import (
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type serverSet struct {
	http chi.Router
	grpc *grpc.Server
}

func newServerSet(opts *serverOpts) *serverSet {
	http := chi.NewMux()
	if len(opts.HTTPMiddlewares) > 0 {
		http.Use(opts.HTTPMiddlewares...)
	}
	http.Mount("/", opts.HTTPMux)
	// TODO check
	http.With(allowCORS)
	srv := &serverSet{
		grpc: grpc.NewServer(opts.GRPCOpts...),
		http: http,
	}
	return srv
}
