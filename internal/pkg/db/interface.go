package db

import (
	"context"
	"io"

	"github.com/go-pg/pg/orm"
)

// DatabaseInterface is a common interface for pg.DB and pg.Tx types
type DatabaseInterface interface {
	Model(model ...interface{}) *orm.Query
	ModelContext(c context.Context, model ...interface{}) *orm.Query
	Select(model interface{}) error
	Insert(model ...interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	ForceDelete(model interface{}) error

	Exec(query interface{}, params ...interface{}) (orm.Result, error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error)
	ExecOne(query interface{}, params ...interface{}) (orm.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error)
	Query(model, query interface{}, params ...interface{}) (orm.Result, error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (orm.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error)

	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error)

	Context() context.Context
	orm.QueryFormatter
}
