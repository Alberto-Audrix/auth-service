package usecases

import (
	"bootcamp-content-interaction-service/domains/sessions"
	"bootcamp-content-interaction-service/domains/sessions/models/requests"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"

	"github.com/google/uuid"
)

type sessionUseCase struct {
	repo sessions.SessionRepository
}

func NewSessionUseCase(repo sessions.SessionRepository) sessions.SessionUseCase {
	return sessionUseCase{repo: repo}
}

func (u sessionUseCase) CreateSession(ctx context.Context, userId uuid.UUID, refreshToken string, isRevoked int) (*sharedresponses.BasicResponse, error) {

	err := u.repo.Create(ctx, userId, refreshToken, 1)

	if err != nil {
		return nil, err
	}

	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Session created",
		},
	}, nil
}

func (uc sessionUseCase) Logout(ctx context.Context, request *requests.LogoutRequest) (*sharedresponses.BasicResponse, error) {

	err := uc.repo.Logout(ctx, request.IsRevoked)

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
