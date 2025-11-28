package interfaces

import (
	"context"
	"zotes/servers/web/db"
	web "zotes/servers/web/db"

	"github.com/jackc/pgx/v5"
)

type CommonContext interface {
	Logger() Logger
	Session() SessionManager
	Sanitizer() Sanitizer
	// S3() S3
}

type DBContext interface {
	SetRLS(r context.Context, tx pgx.Tx) error
	SetSystemRLS(r context.Context, tx pgx.Tx) error
	Client(tx pgx.Tx) *web.Queries
	// Admin(tx pgx.Tx) *admin.Queries

	WithTx(context.Context, int, func(q *db.Queries)) error
	// NewQueries(pgx.Tx) T
	// QueryTx(r context.Context, params StandardSQLCQuery[T], tx pgx.Tx) (any, error)
}

type UserContext interface {
	DBContext
	CommonContext
}
