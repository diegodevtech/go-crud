package controller

import (
	"net/http"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context){

	logger.Info("Initializing DeleteUser Controller Method", zap.String("journey","deleteUser"))
	
	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a valid HEX")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully", zap.String("userId", userId), zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}