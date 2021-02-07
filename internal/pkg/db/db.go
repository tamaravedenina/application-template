package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"google.golang.org/grpc"

	"application-template/internal/pkg/config"
)

type dbLogger struct{}

// GetDB returns connection
func GetDB(ctx context.Context) DatabaseInterface {
	connection := ctx.Value("db").(DatabaseInterface)
	return connection
}

// GetDatabaseConnection returns new pg.DB
func GetDatabaseConnection(cfg config.Database) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Addr,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	if cfg.Options.LogQuery == true {
		db.AddQueryHook(dbLogger{})
	}

	return db
}

// UnaryDatabaseInterceptor set db to context of a unary RPC
func UnaryDatabaseInterceptor(connection DatabaseInterface) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, "db", connection)
		return handler(ctx, req)
	}
}

// BeforeQuery called before queries
func (d dbLogger) BeforeQuery(qe *pg.QueryEvent) {
	if qe.Data == nil {
		qe.Data = make(map[interface{}]interface{})
	}
	qe.Data["queryStartTime"] = time.Now()
}

// AfterQuery called after queries
func (d dbLogger) AfterQuery(qe *pg.QueryEvent) {
	var duration time.Duration
	if qe.Data != nil {
		if v, ok := qe.Data["queryStartTime"]; ok {
			duration = time.Now().Sub(v.(time.Time))
		}
	}
	sql, err := qe.FormattedQuery()
	fmt.Printf("%s: %s, err: %v\n", duration, sql, err)
}
