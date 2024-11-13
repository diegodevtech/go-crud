package service

import (
	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	logger.Info("Initializing DeleteUser Model Method", zap.String("journey", "deleteUser"))
	
	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		logger.Error("Error trying to call deleteUser repository", err, zap.String("journey","deleteUser"))
		return err
	}

	logger.Info("DeleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}