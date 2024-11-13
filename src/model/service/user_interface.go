package service

import (
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(id string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(string) *rest_err.RestErr
}