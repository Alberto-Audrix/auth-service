package repositories

import (
	sessions "bootcamp-content-interaction-service/domains/sessions/entities"
	"bootcamp-content-interaction-service/domains/users"
	"bootcamp-content-interaction-service/domains/users/entities"
	"bootcamp-content-interaction-service/infrastructures"
	"bootcamp-content-interaction-service/shared/util"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type databaseUserRepository struct {
	db         infrastructures.Database
	redisCache *redis.Client
	logger     util.Logger
}

func NewDatabaseUserRepository(db infrastructures.Database, redisClient *redis.Client, logger util.Logger) users.UserRepository {
	return databaseUserRepository{
		db:         db,
		redisCache: redisClient,
		logger:     logger,
	}
}

func (repo databaseUserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	key := "user:" + username

	cached, err := repo.redisCache.Get(ctx, key).Result()
	if err == nil {
		repo.logger.Info("Cache hit - returning users from redis",
			zap.String("cache_key", key),
		)
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return &user, nil
		}
	}

	result := repo.db.GetInstance().WithContext(ctx).Where("username = ?", username).First(&user)
	repo.logger.Info("Get from DB",
		zap.String("username", username),
	)
	if result.Error != nil {
		return nil, result.Error
	}

	userJSON, _ := json.Marshal(user)
	_ = repo.redisCache.Set(ctx, key, userJSON, time.Hour).Err()
	repo.logger.Info("Set user in cache",
		zap.String("user_json", string(userJSON)),
	)

	return &user, nil
}

func (repo databaseUserRepository) FindById(ctx context.Context, id string) (*entities.User, error) {
	var user entities.User
	key := "user:" + id

	cached, err := repo.redisCache.Get(ctx, key).Result()
	if err == nil {
		repo.logger.Info("Cache hit - returning users from redis",
			zap.String("cache_key", key),
		)
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return &user, nil
		}
	}

	result := repo.db.GetInstance().WithContext(ctx).Where("id = ?", id).First(&user)
	repo.logger.Info("Get from DB",
		zap.String("id", id),
	)
	if result.Error != nil {
		return nil, result.Error
	}

	userJSON, _ := json.Marshal(user)
	_ = repo.redisCache.Set(ctx, key, userJSON, time.Hour).Err()
	repo.logger.Info("Set user in cache",
		zap.String("user_json", string(userJSON)),
	)

	return &user, nil
}

func (repo databaseUserRepository) FindSession(ctx context.Context, token string) (*sessions.Session, error) {
	var session sessions.Session
	key := "session:" + token

	cached, err := repo.redisCache.Get(ctx, key).Result()
	if err == nil {
		repo.logger.Info("Cache hit - returning users from redis",
			zap.String("cache_key", key),
		)
		if err := json.Unmarshal([]byte(cached), &session); err == nil {
			return &session, nil
		}
	}

	result := repo.db.GetInstance().WithContext(ctx).Where("refresh_token = ?", token).First(&session)
	repo.logger.Info("Get from DB",
		zap.String("token", token),
	)
	if result.Error != nil {
		return nil, result.Error
	}

	userJSON, _ := json.Marshal(session)
	_ = repo.redisCache.Set(ctx, key, userJSON, time.Hour).Err()
	repo.logger.Info("Set session in cache",
		zap.String("session_json", string(userJSON)),
	)

	return &session, nil
}

func (repo databaseUserRepository) Create(ctx context.Context, name string, username string, email string, password string, bio string, gender string, phone string, country string, profile string) error {

	result := repo.db.GetInstance().WithContext(ctx).Create(
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
		repo.logger.Error("Failed to create user in database",
			zap.Error(result.Error),
		)
		return result.Error
	}

	cacheKey := fmt.Sprintf("user:%s", name)
	if err := repo.redisCache.Set(ctx, cacheKey, "1", time.Hour).Err(); err != nil {
		repo.logger.Error("Failed to create user in Redis cache",
			zap.Error(err),
			zap.String("cache_key", cacheKey),
		)
	}

	repo.logger.Info("User created successfully")
	return nil
}

func (repo databaseUserRepository) CreateSession(ctx context.Context, userId uuid.UUID, refreshToken string, isRevoked int) error {

	result := repo.db.GetInstance().WithContext(ctx).Create(
		&sessions.Session{
			ID:           uuid.New(),
			UserID:       userId,
			RefreshToken: refreshToken,
			IsRevoked:    isRevoked,
			CreatedAt:    time.Now(),
			ExpiredAt:    time.Now().Add(time.Hour),
		},
	)

	if result.Error != nil {
		repo.logger.Error("Failed to create session in database",
			zap.Error(result.Error),
		)
		return result.Error
	}

	cacheKey := fmt.Sprintf("session:%s", refreshToken)
	if err := repo.redisCache.Set(ctx, cacheKey, "1", time.Hour).Err(); err != nil {
		repo.logger.Error("Failed to create session in Redis cache",
			zap.Error(err),
			zap.String("cache_key", cacheKey),
		)
	}

	repo.logger.Info("Session created successfully")
	return nil
}
