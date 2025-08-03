package repos

import (
	"context"
	"user_service/internal/database"
	"user_service/internal/utils"
	"user_service/proto/user"

	"github.com/thanvuc/go-core-lib/log"
)

type userRepo struct {
	logger log.Logger
	sqlc   *database.Queries
}

func (ur *userRepo) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (string, error) {
	id, err := utils.ToUUID(req.UserId)
	if err != nil {
		return "", err
	}
	ur.sqlc.GetUserProfile(ctx, id)
	return "User data for ID: ", nil
}
