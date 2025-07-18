package users

import (
	"bootcamp-content-interaction-service/domains/users/entities"
	"bootcamp-content-interaction-service/domains/users/models/dto/requests"
	"bootcamp-content-interaction-service/domains/users/models/dto/responses"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"

	"github.com/google/uuid"
)

type UserUseCase interface {
	Login(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error)
	GetCurrentUser(ctx context.Context) (*responses.CurrentUserResponse, error)
	SignUp(ctx context.Context, request *requests.SignUpRequest) (*sharedresponses.BasicResponse, error)
}

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	Create(ctx context.Context, name string, username string, email string, password string, bio string, gender string, phone string, country string, profile string) error
	CreateSession(ctx context.Context, userId uuid.UUID, refreshToken string, isRevoked int) error
}
