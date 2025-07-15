package repositories

import (
	"bootcamp-content-interaction-service/domains/sessions"
	"bootcamp-content-interaction-service/domains/sessions/entities"
	"bootcamp-content-interaction-service/infrastructures"
	"context"
	"time"

	"github.com/google/uuid"
)

type databaseSessionRepository struct {
	db infrastructures.Database
}

func NewDatabaseSessionRepository(db infrastructures.Database) sessions.SessionRepository {
	return databaseSessionRepository{
		db: db,
	}
}

func (repo databaseSessionRepository) CreateSession(ctx context.Context, userId uuid.UUID, refreshToken string, isRevoked int64) error {
	result := repo.db.GetInstance().Create(
		&entities.Session{
			ID:           uuid.New(),
			UserID:       userId,
			RefreshToken: refreshToken,
			CreatedAt:    time.Now(),
		},
	)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
