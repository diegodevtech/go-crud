package model

import "github.com/diegodevtech/go-crud/src/configuration/rest_err"

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)
	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:email, 
		password: password, 
	}
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:email, 
		password: password, 
		name: name, 
		age: age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name, 
		age: age,
	}
}