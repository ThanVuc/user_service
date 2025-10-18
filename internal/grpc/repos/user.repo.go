package repos

import (
	"context"
	"fmt"
	"time"
	"user_service/internal/grpc/database"
	"user_service/internal/grpc/utils"
	"user_service/proto/user"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thanvuc/go-core-lib/log"
)

type userRepo struct {
	sqlc   *database.Queries
	logger log.Logger
	pool   *pgxpool.Pool
}

func (ur *userRepo) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*database.GetUserProfileRow, error) {
	userIdUUID, err := utils.ToUUID(req.Id)
	if err != nil {
		return nil, err
	}

	userPr, err := ur.sqlc.GetUserProfile(ctx, userIdUUID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	
	if !userPr.Slug.Valid && userPr.Slug.String == "" {
		var createdAt time.Time
		if userPr.CreatedAt.Valid {
			createdAt = userPr.CreatedAt.Time
		} else {
			createdAt = time.Now()
		}

		newSlug := utils.MakeSlug(userPr.Fullname.String, createdAt)
		_, err := ur.sqlc.UpdateSlugById(ctx, database.UpdateSlugByIdParams{
			ID:   userIdUUID,
			Slug: pgtype.Text{String: newSlug, Valid: true},
		})
		if err != nil {
			return nil, err
		}

		userPr.Slug = pgtype.Text{String: newSlug, Valid: true}
	}

	return &userPr, nil
}

func (ur *userRepo) UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*pgtype.UUID, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("user id is empty")
	}

	userId, err := utils.ToUUID(req.Id)
	if err != nil {
		return nil, err
	}

	userPr, err := ur.sqlc.GetUserProfile(ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if req.Fullname != "" && req.Fullname != userPr.Fullname.String {
		var createdAt time.Time
		if userPr.CreatedAt.Valid {
			createdAt = userPr.CreatedAt.Time
		}
		newSlug := utils.MakeSlug(req.Fullname, createdAt)
		_, err := ur.sqlc.UpdateSlugById(ctx, database.UpdateSlugByIdParams{
			ID:   userId,
			Slug: pgtype.Text{String: newSlug, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}

	var dob pgtype.Timestamptz
	if req.DateOfBirth != 0 {
		dob = pgtype.Timestamptz{
			Time:  time.Unix(req.DateOfBirth, 0),
			Valid: true,
		}
	} else {
		dob = pgtype.Timestamptz{Valid: false}
	}

	userId, err = ur.sqlc.UpdateUserProfile(
		ctx,
		database.UpdateUserProfileParams{
			ID:          userId,
			Fullname:    pgtype.Text{String: req.Fullname, Valid: true},
			Bio:         pgtype.Text{String: req.Bio, Valid: true},
			DateOfBirth: dob,
			Gender:      pgtype.Bool{Bool: req.Gender, Valid: true},
			Sentence:    pgtype.Text{String: req.Sentence, Valid: true},
			Author:      pgtype.Text{String: req.Author, Valid: true},
		},
	)

	if err != nil && err == pgx.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &userId, nil

}
