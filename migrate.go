package bbs

import (
	_ "embed"
	"context"
	"database/sql"
)

//go:embed schema.sql
var sqlSchema string

func MigrateSQL(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlSchema)
	return err
}