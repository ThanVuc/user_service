package repos

import (
	"user_service/global"
	"user_service/internal/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IUserRepo interface {
	GetUsers() ([]database.UserProfile, error)
}

type UserRepo struct {
	sqlc *database.Queries
}

func NewUserRepo() IUserRepo {
	return &UserRepo{
		sqlc: database.New(global.PostgresPool),
	}
}

func (r *UserRepo) GetUsers() ([]database.UserProfile, error) {
	users := make([]database.UserProfile, 0)
	users = append(users, database.UserProfile{
		UserID:   pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Username: "testuser",
		Fullname: "Test User",
	})
	return users, nil
}
