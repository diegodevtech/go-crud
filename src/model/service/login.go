package service

import (
	"fmt"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initializing Login Model Method", zap.String("journey", "login"))

	fmt.Println("=================")

	fmt.Println(userDomain.GetEmail(), userDomain.GetPassword())
	// userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	
	fmt.Println("=================")

	user, err := ud.findUserByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, err
	}

	logger.Info("Login service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "login"),
	)

	return user, nil
}