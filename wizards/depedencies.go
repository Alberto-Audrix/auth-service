package wizards

import (
	"bootcamp-content-interaction-service/config"
	usersHttp "bootcamp-content-interaction-service/domains/users/handlers/http"
	usersRepo "bootcamp-content-interaction-service/domains/users/repositories"
	usersUc "bootcamp-content-interaction-service/domains/users/usecases"
	"bootcamp-content-interaction-service/infrastructures"
)

var (
	Config           = config.GetConfig()
	PostgresDatabase = infrastructures.NewPostgresDatabase(Config)
	UserDatabaseRepo = usersRepo.NewDatabaseUserRepository(PostgresDatabase)
	UserUc           = usersUc.NewUserUseCase(UserDatabaseRepo)
	UserHttp         = usersHttp.NewUserHttp(UserUc)
)
