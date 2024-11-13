package service

import (
	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Initializing UpdateUser Model Method", zap.String("journey", "updateUser"))
	
	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		logger.Error("Error trying to call updateUser repository", err, zap.String("journey","updateUser"))
		return err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}