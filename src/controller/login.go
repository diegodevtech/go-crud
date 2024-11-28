package controller

import (
	"net/http"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/validation"
	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) Login(c *gin.Context) {
	logger.Info("Initializing Login Controller Method", zap.String("journey", "login"))

	var loginRequest request.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		logger.Error("Error trying to validate login info", err, zap.String("journey", "login"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewLoginDomain(
		loginRequest.Email,
		loginRequest.Password,
	)

	domainResult, err := uc.service.LoginService(domain)
	if err != nil {
		logger.Error("Error trying to call login service", err, zap.String("journey","login"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Login controller executed successfully", zap.String("userId", domainResult.GetID()), zap.String("journey","login"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}