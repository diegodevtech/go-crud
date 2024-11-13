package service

import (
	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initializing CreateUser Model Method", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error trying to call createUser repository", err, zap.String("journey","createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}