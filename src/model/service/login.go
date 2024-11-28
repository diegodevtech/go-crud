package service

import (

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Initializing Login Model Method", zap.String("journey", "login"))

	// userDomain.EncryptPassword() // md5 only

	user, err := ud.findUserByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("Login service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "login"),
	)

	return user, token, nil
}