package usecases

import (
	"bootcamp-content-interaction-service/domains/sessions"
	"bootcamp-content-interaction-service/domains/sessions/models/requests"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"
	"time"
)

type sessionUseCase struct {
	repo sessions.SessionRepository
}

func NewSessionUseCase(repo sessions.SessionRepository) sessions.SessionUseCase {
	return sessionUseCase{repo: repo}
}

func (uc sessionUseCase) Logout(ctx context.Context, request *requests.LogoutRequest, token string) (*sharedresponses.BasicResponse, error) {

	session, err := uc.repo.FindSession(ctx, token)

	if err != nil {
		return nil, err
	}

	session.IsRevoked = request.IsRevoked
	session.ExpiredAt = time.Now()

	_, err = uc.repo.Logout(ctx, session, token)

	if err != nil {
		return nil, err
	}

	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Logout Success",
		},
	}, nil
}
