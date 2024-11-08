package controller

import (
	"net/http"

	// "github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/validation"
	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("Initializing CreateUser Controller Method",
		zap.String("journey","createUser"),
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

	response := response.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey","createUser"),
	)

	c.JSON(http.StatusOK, response)
}
