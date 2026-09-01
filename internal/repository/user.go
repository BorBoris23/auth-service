package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	dbConn *pgx.Conn
}

type User struct {
	ID           int
	Login        string
	PasswordHash string
	Role         Role
}

type Role struct {
	ID   int
	Name string
}

func NewUserRepository(dbConn *pgx.Conn) *UserRepository {
	return &UserRepository{
		dbConn: dbConn,
	}
}

func (r *UserRepository) FindUser(
	ctx context.Context,
	login string,
) (*User, error) {
	var user User

	err := r.dbConn.QueryRow(
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

func (r *UserRepository) FindRoleByID(
	ctx context.Context,
	roleID int,
) (*Role, error) {
	var role Role

	err := r.dbConn.QueryRow(
		ctx,
		`
			SELECT id, name
			FROM roles
			WHERE id = $1
		`,
		roleID,
	).Scan(
		&role.ID,
		&role.Name,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}
