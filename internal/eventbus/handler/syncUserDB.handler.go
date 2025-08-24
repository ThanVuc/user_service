package handler

import (
	"context"
	"encoding/json"
	"time"
	"user_service/global"
	"user_service/internal/grpc/database"
	"user_service/internal/grpc/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thanvuc/go-core-lib/log"
)

type SyncAuthHandler struct {
	logger log.Logger
	sqlc   *database.Queries
}

type UserOutboxPayload struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

func NewSyncAuthHandler() *SyncAuthHandler {
	return &SyncAuthHandler{
		logger: global.Logger,
		sqlc:   database.New(global.PostgresPool),
	}
}

func (h *SyncAuthHandler) SyncUserDB(ctx context.Context, payload []byte) error {
	var userPayload UserOutboxPayload
	if err := json.Unmarshal(payload, &userPayload); err != nil {
		return err
	}

	// Insert user into the database
	userId, err := utils.ToUUID(userPayload.UserID)
	if err != nil {
		return err
	}

	createdAt := pgtype.Timestamptz{
		Time:  time.Unix(userPayload.CreatedAt, 0),
		Valid: true,
	}

	_, err = h.sqlc.InsertUser(ctx, database.InsertUserParams{
		ID:        userId,
		Email:     userPayload.Email,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	})
	if err != nil {
		return err
	}

	return nil
}
