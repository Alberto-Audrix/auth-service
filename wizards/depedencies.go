package wizards

import (
	"bootcamp-content-interaction-service/config"
	sessionsHttp "bootcamp-content-interaction-service/domains/sessions/handlers/http"
	sessionsRepo "bootcamp-content-interaction-service/domains/sessions/repositories"
	sessionsUc "bootcamp-content-interaction-service/domains/sessions/usecases"
	usersHttp "bootcamp-content-interaction-service/domains/users/handlers/http"
	usersRepo "bootcamp-content-interaction-service/domains/users/repositories"
	usersUc "bootcamp-content-interaction-service/domains/users/usecases"
	"bootcamp-content-interaction-service/infrastructures"
	"bootcamp-content-interaction-service/shared/util"
)

var (
	Config              = config.GetConfig()
	PostgresDatabase    = infrastructures.NewPostgresDatabase(Config)
	UserDatabaseRepo    = usersRepo.NewDatabaseUserRepository(PostgresDatabase, RedisClient, LoggerInstance)
	UserUc              = usersUc.NewUserUseCase(UserDatabaseRepo)
	UserHttp            = usersHttp.NewUserHttp(UserUc)
	SessionDatabaseRepo = sessionsRepo.NewDatabaseSessionRepository(PostgresDatabase, RedisClient, LoggerInstance)
	SessionUc           = sessionsUc.NewSessionUseCase(SessionDatabaseRepo)
	SessionHttp         = sessionsHttp.NewSessionHttp(SessionUc)
	RedisClient         = infrastructures.InitRedis()
	LoggerInstance, _   = util.NewLogger()
)
