package seed

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SeedRoles(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
		INSERT INTO roles (name)
		VALUES
			('user'),
			('moderator'),
			('editor'),
			('manager'),
			('admin')
		ON CONFLICT (name) DO NOTHING
	`)

	return err
}
