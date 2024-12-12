package controller

import (
	// "fmt"
	"net/http"
	"net/mail"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// FindUserByID retrieves user information based on the provided user ID.
// @Summary Find User by ID
// @Description Retrieves user details based on the user ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: User ID is not a valid ID"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserById/{userId} [get]
func (uc *userControllerInterface) FindUserByID(c *gin.Context){
	logger.Info("Initializing FindUserByID Controller Method", zap.String("journey", "findUserByID"))

	// user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	// if err != nil {
	// 	c.JSON(err.Code, err)
	// 	return
	// }

	// logger.Info(fmt.Sprintf("User authenticated: %#v", user))
	
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

// FindUserByEmail retrieves user information based on the provided email.
// @Summary Find User by Email
// @Description Retrieves user details based on the email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userEmail path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: UserEmail is not a valid email"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserByEmail/{userEmail} [get]
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context){
	logger.Info("Initializing FindUserByEmail Controller Method", zap.String("journey", "findUserByEmail"))

	// user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	// if err != nil {
	// 	c.JSON(err.Code, err)
	// 	return
	// }

	// logger.Info(fmt.Sprintf("User authenticated: %#v", user))

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