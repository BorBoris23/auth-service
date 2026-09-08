package grpc

import (
	"context"

	"auth-service/internal/services"

	usershpb "github.com/BorBoris23/auth-proto/gen/users"
)

type UserServer struct {
	usershpb.UnimplementedUserServiceServer

	userService *services.UserService
}

func NewUserServer(userService *services.UserService) *UserServer {
	return &UserServer{
		userService: userService,
	}
}

func (s *UserServer) ValidateUserId(
	ctx context.Context,
	req *usershpb.ValidateUsersRequest,
) (*usershpb.ValidateUsersResponse, error) {
	result, err := s.userService.ValidateUserId(ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	return &usershpb.ValidateUsersResponse{
		Users: result,
	}, nil
}
