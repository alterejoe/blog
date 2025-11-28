package routes

import (
	"context"
	"fmt"
	"zotes/servers/web/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (d *Dependencies[T]) GetDBPool() *pgxpool.Pool {
	return d.db
}

func (d *Dependencies[T]) SetRLS(r context.Context, tx pgx.Tx) error {
	id := d.Session().Get(r, "authenticatedUserID")
	username := d.Session().Get(r, "user_name")
	useremail := d.Session().Get(r, "user_email")

	if id == nil {
		return fmt.Errorf("missing authenticated user id in session")
	}

	if _, err := tx.Exec(r,
		`SELECT
        set_config('app.user_id', $1::text, true),
        set_config('app.user_name', $2::text, true),
        set_config('app.user_email', $3::text, true);`,
		id, username, useremail,
	); err != nil {
		return fmt.Errorf("set RLS failed: %w", err)
	}

	return nil
}

func (d *Dependencies[T]) SetSystemRLS(r context.Context, tx pgx.Tx) error {
	id := "00000000-0000-0000-0000-000000000000"
	username := "SYSTEM"
	useremail := "system@harpky.com"

	if _, err := tx.Exec(r,
		`SELECT
        set_config('app.user_id', $1::text, true),
        set_config('app.user_name', $2::text, true),
        set_config('app.user_email', $3::text, true);`,
		id, username, useremail,
	); err != nil {
		return fmt.Errorf("set RLS failed: %w", err)
	}

	return nil
}

type RLSMode int

const (
	RLSUser RLSMode = iota
	RLSSystem
	RLSNone
)

func (d *Dependencies[T]) WithTx(
	ctx context.Context,
	mode RLSMode,
	fn func(*db.Queries) error,
) error {
	tx, err := d.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	switch mode {
	case RLSUser:
		if err := d.SetRLS(ctx, tx); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	case RLSSystem:
		if err := d.SetSystemRLS(ctx, tx); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	case RLSNone:
		// no-op
	}

	q := db.New(tx)

	if err := fn(q); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
