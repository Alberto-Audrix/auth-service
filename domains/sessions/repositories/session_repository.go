package repositories

import (
	"bootcamp-content-interaction-service/domains/sessions"
	"bootcamp-content-interaction-service/domains/sessions/entities"
	"bootcamp-content-interaction-service/infrastructures"
	"bootcamp-content-interaction-service/shared/util"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type databaseSessionRepository struct {
	db         infrastructures.Database
	redisCache *redis.Client
	logger     util.Logger
}

func NewDatabaseSessionRepository(db infrastructures.Database, redisClient *redis.Client, logger util.Logger) sessions.SessionRepository {
	return databaseSessionRepository{
		db:         db,
		redisCache: redisClient,
		logger:     logger,
	}
}

func (repo databaseSessionRepository) Logout(ctx context.Context, session *entities.Session, token string) (*entities.Session, error) {
	repo.logger.Info("Attempting to logout")

	result := repo.db.GetInstance().WithContext(ctx).Save(session)
	if result.Error != nil {
		return nil, result.Error
	}

	postJSON, err := json.Marshal(session)
	if err == nil {
		key := "post:" + session.ID.String()
		_ = repo.redisCache.Set(ctx, key, postJSON, time.Hour).Err()
		repo.logger.Info("Update session in redis",
			zap.String("post_id", session.ID.String()),
		)
	}

	return session, nil
}

func (repo databaseSessionRepository) FindSession(ctx context.Context, token string) (*entities.Session, error) {
	var session entities.Session
	var logger = zap.NewExample()
	key := "session:" + token

	cached, err := repo.redisCache.Get(ctx, key).Result()
	if err == nil {
		logger.Info("Redis HIT: " + key)
		if err := json.Unmarshal([]byte(cached), &session); err == nil {
			return &session, nil
		}
	}

	result := repo.db.GetInstance().WithContext(ctx).Where("refresh_token = ?", token).First(&session)
	logger.Info("Get from DB with refresh token" + token)
	if result.Error != nil {
		return nil, result.Error
	}

	postJSON, _ := json.Marshal(session)
	_ = repo.redisCache.Set(ctx, key, postJSON, time.Hour).Err()
	logger.Info("Set session in redis cache " + string(postJSON))

	return &session, nil
}
