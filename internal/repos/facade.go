package repos

import (
	"context"
	"user_service/global"
	"user_service/internal/database"
	"user_service/proto/user"
)

type (
	UserRepo interface {
		GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (string, error)
	}
)

func NewUserRepo() UserRepo {
	return &userRepo{
		logger: global.Logger,
		sqlc:   database.New(global.PostgresPool),
	}
}
