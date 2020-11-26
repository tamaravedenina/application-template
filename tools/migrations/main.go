// This is custom goose binary with sqlite3 support only.

package main

import (
	"application-template/internal/pkg/config"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	_ "application-template/migrations"
	_ "github.com/mattn/go-sqlite3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
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
	}

	// init db
	connConfig, _ := pgx.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		appConfig.Database.User,
		appConfig.Database.Password,
		appConfig.Database.Addr,
		appConfig.Database.Database,
	))
	connStr := stdlib.RegisterConnConfig(connConfig)
	dbConnect, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		if err := dbConnect.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	goose.SetTableName("migration")

	flags.Parse(os.Args[1:])
	args := flags.Args()

	command := args[0]
	if len(args) < 1 {
		flags.Usage()
		return
	}

	arguments := make([]string, 0)
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, dbConnect.DB, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
