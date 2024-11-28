package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context){
	logger.Info("Initializing FindUserByID Controller Method", zap.String("journey", "findUserByID"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))
	
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey","findUserByID"))
		errorMessage := rest_err.NewBadRequestError("User ID is not a valid ID")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID service", err, zap.String("journey","findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context){
	logger.Info("Initializing FindUserByEmail Controller Method", zap.String("journey", "findUserByEmail"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey","findUserByEmail"))

		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)

	if err != nil {
		logger.Error("Error trying to call findUserByEmail service", err, zap.String("journey","findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("journey", "findUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}