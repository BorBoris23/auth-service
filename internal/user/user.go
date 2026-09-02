package user

import (
	"context"
	"errors"

	"auth-service/internal/role"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

type User struct {
	ID           int
	Login        string
	PasswordHash string
	Role         role.Role
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUser(
	ctx context.Context,
	login string,
) (*User, error) {
	var user User

	err := r.db.QueryRow(
		ctx,
		`
			SELECT id, login, password, role_id
			FROM users
			WHERE login = $1
		`,
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.PasswordHash,
		&user.Role.ID,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("invalid credentials")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
