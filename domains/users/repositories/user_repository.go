package repositories

import (
	"bootcamp-content-interaction-service/domains/users"
	"bootcamp-content-interaction-service/domains/users/entities"
	"bootcamp-content-interaction-service/infrastructures"
	"context"
	"time"

	"github.com/google/uuid"
)

type databaseUserRepository struct {
	db infrastructures.Database
}

func NewDatabaseUserRepository(db infrastructures.Database) users.UserRepository {
	return databaseUserRepository{
		db: db,
	}
}

func (repo databaseUserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	result := repo.db.GetInstance().Preload("Wallet").First(&user, "username = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo databaseUserRepository) Create(ctx context.Context, name string, username string, email string, password string, bio string, gender string, phone string, country string, profile string) error {
	result := repo.db.GetInstance().Create(
		&entities.User{
			ID:        uuid.New(),
			Name:      name,
			Username:  username,
			Email:     email,
			Password:  password,
			Bio:       bio,
			Gender:    gender,
			Phone:     phone,
			Country:   country,
			Profile:   profile,
			CreatedAt: time.Now(),
		},
	)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
