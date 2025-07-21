package users

import (
	session "bootcamp-content-interaction-service/domains/sessions/entities"
	"bootcamp-content-interaction-service/domains/users/entities"
	"bootcamp-content-interaction-service/domains/users/models/dto/requests"
	"bootcamp-content-interaction-service/domains/users/models/dto/responses"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"

	"github.com/google/uuid"
)

type UserUseCase interface {
	Login(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error)
	GetCurrentUser(ctx context.Context, token string) (*responses.CurrentUserResponse, error)
	SignUp(ctx context.Context, request *requests.SignUpRequest) (*sharedresponses.BasicResponse, error)
}

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	FindById(ctx context.Context, id string) (*entities.User, error)
	FindSession(ctx context.Context, token string) (*session.Session, error)
	Create(ctx context.Context, name string, username string, email string, password string, bio string, gender string, phone string, country string, profile string) error
	CreateSession(ctx context.Context, userId uuid.UUID, refreshToken string, isRevoked int) error
}
