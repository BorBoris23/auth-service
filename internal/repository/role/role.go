package role

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepository struct {
	db *pgxpool.Pool
}

type Role struct {
	ID   int
	Name string
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) FindByID(
	ctx context.Context,
	roleID int,
) (*Role, error) {
	var role Role

	err := r.db.QueryRow(
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
