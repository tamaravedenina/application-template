package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	_ "github.com/utrack/clay/doc/example/static/statik"
	"github.com/utrack/clay/v2/log"
	"github.com/utrack/clay/v2/transport/middlewares/mwgrpc"

	um "application-template/internal/app/user"
	"application-template/internal/pkg/config"
	"application-template/internal/pkg/db"
	"application-template/internal/pkg/server"
)

func main() {
	err := mainWithoutFatal()
	if err != nil {
		logrus.Fatal(err)
	}
}

func mainWithoutFatal() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	appConfig, err := config.InitConfigFromFile(fmt.Sprintf("./config/%s.yml", env))
	if err != nil {
		return err
	}

	fmt.Printf("appConfig: %+v\n", appConfig)

	databaseConnection := db.GetDatabaseConnection(appConfig.Database)
	defer func() {
		if closeErr := databaseConnection.Close(); closeErr != nil {
			err = closeErr
			return
		}
	}()

	staticFS, err := fs.New()
	if err != nil {
		return err
	}

	hmux := chi.NewRouter()
	hmux.Mount("/", http.FileServer(staticFS))

	userModule := um.BuildUserModule()

	srv := server.NewServer(
		appConfig.GrpcPort,
		appConfig.HttpPort,
		// Pass out mux with Swagger UI
		server.WithHTTPMux(hmux),
		// Recover from both HTTP and gRPC panics and use our own middleware
		server.WithGRPCUnaryMiddlewares(
			mwgrpc.UnaryPanicHandler(log.Default),
			db.UnaryDatabaseInterceptor(databaseConnection),
		),
	)

	err = srv.Run(userModule)

	return err
}
