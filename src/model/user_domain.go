package model

import (
	// "crypto/md5"
	// "encoding/hex"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(
	email, password, name string, 
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncrtpyPassword() {
	// hash := md5.New()
	// defer hash.Reset()
	// hash.Write([]byte(ud.Password))
	// ud.Password = hex.EncodeToString(hash.Sum(nil))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Fail attempting to encrypt password", err)
	}
	ud.Password = string(hashedPassword)
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
