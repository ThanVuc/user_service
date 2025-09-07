package repos

import (
	"context"
	"user_service/global"
	"user_service/internal/grpc/database"
	"user_service/proto/user"

	"github.com/jackc/pgx/v5/pgtype"
)

type (
	UserRepo interface {
		GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*[]database.GetUserProfileRow, error)
		UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*pgtype.UUID, error)
	}
)


func NewUserRepo() UserRepo {
	return &userRepo{
		logger: global.Logger,
		sqlc:   database.New(global.PostgresPool),
		pool:   global.PostgresPool,
	}
}

