package wizards

import (
	"bootcamp-content-interaction-service/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterServer(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/login", UserHttp.Login)
		auth.POST("/signup", UserHttp.SignUp)
		auth.PUT("/logout", SessionHttp.Logout)

		auth.Use(middlewares.AuthMiddleware())
		auth.GET("/me", UserHttp.GetCurrentUser)
	}
}
