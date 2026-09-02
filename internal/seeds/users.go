package seed

import (
	"context"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(ctx context.Context, conn *pgxpool.Pool) error {
	var adminRoleID int

	err := conn.QueryRow(
		ctx,
		`SELECT id FROM roles WHERE name = 'admin'`,
	).Scan(&adminRoleID)
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		name := faker.Name()
		login := fmt.Sprintf("user_%d", i+1)
		roleID := (i % 4) + 1
		password := "123456"

		if i == 0 {
			name = "Admin"
			login = "admin"
			roleID = adminRoleID
		}

		passwordHash, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}

		_, err = conn.Exec(ctx, `
			INSERT INTO users (name, login, password, role_id)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (login) DO NOTHING
		`, name, login, string(passwordHash), roleID)

		if err != nil {
			return err
		}
	}

	return nil
}
