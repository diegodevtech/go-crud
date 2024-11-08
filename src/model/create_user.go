package model

import (
	"fmt"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Initializing CreateUser Model Method", zap.String("journey", "createUser"))

	ud.EncrtpyPassword()
	fmt.Println(ud)

	return nil
}