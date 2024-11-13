package controller

import (
	"net/http"
	"strings"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/validation"
	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context){
	logger.Info("Initializing UpdateUser Controller Method",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")

	err := c.ShouldBindJSON(&userRequest)

	if err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"),
		)
		// rest_error := rest_err.NewBadRequestError(fmt.Sprintf("There are some invalid fields. Error: %s", err.Error()))
		rest_error := validation.ValidateUserError(err)
		c.JSON(rest_error.Code, rest_error)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err2 := uc.service.UpdateUserService(userId, domain)
	if err2 != nil {
		logger.Error("Error trying to call updateUser service",err, zap.String("journey", "updateUser"))
		c.JSON(err2.Code, err)
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "createUser"),
	)

	c.Status(http.StatusOK)
}