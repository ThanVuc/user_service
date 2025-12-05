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
	"go.uber.org/zap"
)

type SyncAuthHandler struct {
	logger log.Logger
	sqlc   *database.Queries
}

type UserOutboxPayload struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	Fullname  string `json:"name"`
	Picture   string `json:"avatar_url"`
}

func NewSyncAuthHandler() *SyncAuthHandler {
	return &SyncAuthHandler{
		logger: global.Logger,
		sqlc:   database.New(global.PostgresPool),
	}
}

func (h *SyncAuthHandler) SyncUserDB(ctx context.Context, payload []byte) error {
	requestID := utils.GetRequestIDFromOutgoingContext(ctx)
	var userPayload UserOutboxPayload
	if err := json.Unmarshal(payload, &userPayload); err != nil {
		return err
	}
	h.logger.Info("Sync user DB handler invoked", requestID, zap.String("picture", userPayload.Picture))
	
	if userPayload.Email == "" && userPayload.Fullname == "" && userPayload.CreatedAt == 0 {
		userId, err := utils.ToUUID(userPayload.UserID)
		if err != nil {
			h.logger.Error("Failed to parse user ID", requestID, zap.Error(err))
			return err
		}

		_, err = h.sqlc.UpdateAvatarById(ctx, database.UpdateAvatarByIdParams{
			ID:        userId,
			AvatarUrl: pgtype.Text{String: userPayload.Picture, Valid: userPayload.Picture != ""},
		})
		if err != nil {
			h.logger.Error("Failed to update user avatar", requestID, zap.Error(err))
			return err
		}

		return nil
	}

	userId, err := utils.ToUUID(userPayload.UserID)
	if err != nil {
		h.logger.Error("Failed to parse user ID", requestID, zap.Error(err))
		return err
	}

	createdAt := pgtype.Timestamptz{
		Time:  time.Unix(userPayload.CreatedAt, 0),
		Valid: true,
	}

	_, err = h.sqlc.InsertUser(ctx, database.InsertUserParams{
		ID:        userId,
		Column2:   userPayload.Email,
		Column3:   pgtype.Text{String: userPayload.Fullname, Valid: userPayload.Fullname != ""},
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	})
	if err != nil {
		h.logger.Error("Failed to insert user", requestID, zap.Error(err))
		return err
	}

	return nil
}
