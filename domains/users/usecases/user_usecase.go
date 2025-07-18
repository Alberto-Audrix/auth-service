package usecases

import (
	"bootcamp-content-interaction-service/domains/users"
	"bootcamp-content-interaction-service/domains/users/models/dto"
	"bootcamp-content-interaction-service/domains/users/models/dto/requests"
	"bootcamp-content-interaction-service/domains/users/models/dto/responses"
	"bootcamp-content-interaction-service/shared/constant"
	sharedresponses "bootcamp-content-interaction-service/shared/models/responses"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo users.UserRepository
}

func NewUserUseCase(repo users.UserRepository) users.UserUseCase {
	return userUseCase{repo: repo}
}

func (u userUseCase) Login(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error) {
	user, err := u.repo.FindByUsername(ctx, request.Username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(constant.JWT_SECRET)

	if err != nil {
		return nil, err
	}

	revoked := 1
	err = u.repo.CreateSession(ctx, user.ID, tokenString, revoked)

	if err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		AccessToken: tokenString,
	}, nil
}

func (uc userUseCase) GetCurrentUser(ctx context.Context) (*responses.CurrentUserResponse, error) {
	authUserDto := ctx.Value("user").(*dto.AuthUserDto)

	user, err := uc.repo.FindByUsername(ctx, authUserDto.Name)

	if err != nil {
		return nil, err
	}

	return &responses.CurrentUserResponse{
		Id:       user.ID.String(),
		Name:     user.Name,
		Username: user.Username,
	}, nil
}

func (u userUseCase) SignUp(ctx context.Context, request *requests.SignUpRequest) (*sharedresponses.BasicResponse, error) {
	if request.Password != request.ConfirmPassword {
		return nil, errors.New("password must be same with confirm password")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		return nil, errors.New("error when hashing password")
	}

	err = u.repo.Create(ctx, request.Name, request.Username, request.Email, string(b), request.Bio, request.Gender, request.Phone, request.Country, request.Profile)

	if err != nil {
		return nil, err
	}

	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Account created",
		},
	}, nil
}
