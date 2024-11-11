package model

import (
	// "crypto/md5"
	// "encoding/hex"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string, 
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	// hash := md5.New()
	// defer hash.Reset()
	// hash.Write([]byte(ud.Password))
	// ud.Password = hex.EncodeToString(hash.Sum(nil))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Fail attempting to encrypt password", err)
	}
	ud.password = string(hashedPassword)
}
