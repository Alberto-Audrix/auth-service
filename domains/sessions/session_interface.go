package sessions

import (
	"bootcamp-content-interaction-service/domains/sessions/models/requests"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"

	"github.com/google/uuid"
)

type SessionUseCase interface {
	Logout(ctx context.Context, request *requests.LogoutRequest) (*sharedresponses.BasicResponse, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, userId uuid.UUID, refreshToken string, isRevoked int64) error
}
