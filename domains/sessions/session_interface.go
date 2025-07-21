package sessions

import (
	"bootcamp-content-interaction-service/domains/sessions/entities"
	"bootcamp-content-interaction-service/domains/sessions/models/requests"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"
)

type SessionUseCase interface {
	Logout(ctx context.Context, request *requests.LogoutRequest, token string) (*sharedresponses.BasicResponse, error)
}

type SessionRepository interface {
	Logout(ctx context.Context, session *entities.Session, token string) (*entities.Session, error)
	FindSession(ctx context.Context, token string) (*entities.Session, error)
}
