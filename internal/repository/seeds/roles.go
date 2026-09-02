package seed

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedRoles(ctx context.Context, conn *pgxpool.Pool) error {
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
