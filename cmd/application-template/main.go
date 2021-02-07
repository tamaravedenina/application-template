package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/utrack/clay/v2/log"
	"github.com/utrack/clay/v2/transport/middlewares/mwgrpc"

	um "application-template/internal/app/user"
	"application-template/internal/pkg/config"
	"application-template/internal/pkg/db"
	"application-template/internal/pkg/server"
)

func main() {
	// get env
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	// init config
	appConfig, err := config.InitConfigFromFile(fmt.Sprintf("./config/%s.yml", env))
	if err != nil {
		logrus.Fatal(err)
	} else {
		// print cfg to log
		fmt.Printf("appConfig: %+v\n", appConfig)
	}

	// init database
	databaseConnection := db.GetDatabaseConnection(appConfig.Database)
	defer func() {
		if err := databaseConnection.Close(); err != nil {
			logrus.Fatalf("failed to close DB: %v\n", err)
		}
	}()

	// Wire up our bundled Swagger UI
	staticFS, err := fs.New()
	if err != nil {
		logrus.Fatal(err)
	}
	hmux := chi.NewRouter()
	hmux.Mount("/", http.FileServer(staticFS))


	//create user module
	userModule := um.BuildUserModule()

	// create server
	srv := server.NewServer(
		appConfig.GrpcPort,
		appConfig.HttpPort,
		// Pass our mux with Swagger UI
		server.WithHTTPMux(hmux),
		// Recover from both HTTP and gRPC panics and use our own middleware
		server.WithGRPCUnaryMiddlewares(
			mwgrpc.UnaryPanicHandler(log.Default),
			db.UnaryDatabaseInterceptor(databaseConnection),
		),
	)

	// run server
	err = srv.Run(userModule)
	if err != nil {
		logrus.Fatal(err)
	}
}
