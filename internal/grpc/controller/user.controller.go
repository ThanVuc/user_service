package controller

import (
	"context"
	"user_service/internal/grpc/services"
	"user_service/internal/grpc/utils"
	"user_service/proto/common"
	"user_service/proto/user"
)

type UserController struct {
	user.UnimplementedUserServiceServer
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}


func (uc *UserController) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error) {
	return utils.WithSafePanic(ctx, req, uc.userService.GetUserProfile)
}

func (uc *UserController) UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*common.EmptyResponse, error){
	return utils.WithSafePanic(ctx, req, uc.userService.UpdateUserProfile)
}