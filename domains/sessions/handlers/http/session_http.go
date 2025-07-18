package http

import (
	"bootcamp-content-interaction-service/domains/sessions"
	"bootcamp-content-interaction-service/domains/sessions/models/requests"
	"bootcamp-content-interaction-service/shared/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type SessionHttp struct {
	uc sessions.SessionUseCase
}

func NewSessionHttp(uc sessions.SessionUseCase) *SessionHttp {
	return &SessionHttp{uc: uc}
}

func (handler *SessionHttp) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &requests.LogoutRequest{}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.StructCtx(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	result, err := handler.uc.Logout(ctx, requestBody)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	c.Header("Authorization", "")

	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: result,
	})
}
