package user

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"auth-service/internal/repository/role"
)

type UserRepository struct {
	db *pgxpool.Pool
}

type User struct {
	ID           int
	Name         string
	Login        string
	PasswordHash string
	Role         role.Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	login string,
	passwordHash string,
) (*User, error) {
	var user User

	err := r.db.QueryRow(
		ctx,
		`
			INSERT INTO users (
				name,
				login,
				password,
				role_id
			)
			VALUES ($1, $2, $3, 1)
			RETURNING
				id,
				name,
				login,
				password,
				role_id,
				created_at,
				updated_at
		`,
		name,
		login,
		passwordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.PasswordHash,
		&user.Role.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
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
