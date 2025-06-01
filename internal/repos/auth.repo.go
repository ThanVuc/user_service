package repos

import (
	"context"
	"fmt"
	"user_service/global"
	"user_service/internal/database"
	"user_service/pkg/response"

	"github.com/jackc/pgx/v5/pgtype"
)

type IAuthRepo interface {
	CreateAuthor(c context.Context) bool
	CreateUserAndAuthor(c context.Context) bool
}

type AuthRepo struct {
	sqlc *database.Queries
}

func NewAuthRepo() IAuthRepo {
	return &AuthRepo{
		sqlc: database.New(global.PostgresPool),
	}
}

// All the below methods are for testing purposes only
func (ur *AuthRepo) CreateAuthor(c context.Context) bool {
	ur.sqlc.CreateAuthor(c, database.CreateAuthorParams{
		Name: "John Doe",
		Bio:  pgtype.Text{String: "A sample bio", Valid: true},
	})
	fmt.Println("Author created successfully")
	return true
}

func (ur *AuthRepo) CreateUserAndAuthor(c context.Context) bool {
	err := ur.sqlc.CreateUserAndAuthor(c, database.CreateUserAndAuthorParams{
		Name:    "Jane Doe",
		Age:     pgtype.Int4{Int32: 30, Valid: true},
		Gender:  pgtype.Text{String: "male", Valid: true},
		Address: pgtype.Text{String: "123 Main St", Valid: true},
		Name_2:  "Jane Author",
		Bio:     pgtype.Text{String: "An author bio", Valid: true},
	})
	if err != nil {
		fmt.Println("Failed to create user and author: logggggg")
		panic(response.InternalServerError(fmt.Sprintf("Failed to create user and author: %v", err)))
	}
	return true
}
