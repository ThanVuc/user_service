package services

import (
	"context"
	"user_service/internal/repos"
	"user_service/proto/user"
)

type userService struct {
	userRepo repos.UserRepo
}

func (us *userService) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error) {
	us.userRepo.GetUserProfile(ctx, req)
	return &user.GetUserProfileResponse{}, nil
}
