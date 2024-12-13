package controller

import (
	"net/http"

	// "github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/validation"
	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/view"

	// "github.com/diegodevtech/go-crud/src/controller/model/response"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
)

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequest true "User information for registration"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createUser [post]
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Initializing CreateUser Controller Method",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		// rest_error := rest_err.NewBadRequestError(fmt.Sprintf("There are some invalid fields. Error: %s", err.Error()))
		rest_error := validation.ValidateUserError(err)
		c.JSON(rest_error.Code, rest_error)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err2 := uc.service.CreateUserService(domain)
	if err2 != nil {
		logger.Error("Error trying to call createUser service",err2, zap.String("journey", "createUser"))
		c.JSON(err2.Code, err2)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
