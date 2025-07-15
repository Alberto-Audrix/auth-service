package http

import (
	"bootcamp-content-interaction-service/domains/users"
	"bootcamp-content-interaction-service/domains/users/models/dto/requests"
	"bootcamp-content-interaction-service/shared/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserHttp struct {
	uc users.UserUseCase
}

func NewUserHttp(uc users.UserUseCase) *UserHttp {
	return &UserHttp{uc: uc}
}

func (handler *UserHttp) Login(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &requests.LoginRequest{}

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

	result, err := handler.uc.Login(ctx, requestBody)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: result,
	})
}

func (handler *UserHttp) GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := handler.uc.GetCurrentUser(ctx)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: res,
	})
}

func (handler *UserHttp) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &requests.SignUpRequest{}

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

	result, err := handler.uc.SignUp(ctx, requestBody)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, result)
}
