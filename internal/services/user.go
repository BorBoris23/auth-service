package services

import (
	"auth-service/internal/repository/role"
	"auth-service/internal/repository/user"
	"context"
)

type UserService struct {
	userRepository *user.UserRepository
	roleRepository *role.RoleRepository
}

func NewUserService(
	userRepository *user.UserRepository,
	roleRepository *role.RoleRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (s *UserService) ValidateUserId(
	ctx context.Context,
	userIds []int64,
) (map[int64]bool, error) {
	result := make(map[int64]bool)

	for _, id := range userIds {
		exists, err := s.userRepository.FindUserById(ctx, id)
		if err != nil {
			return nil, err
		}

		result[id] = exists
	}

	return result, nil
}
