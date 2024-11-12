package model

import (
	// "crypto/md5"
	// "encoding/hex"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetID() string {
	return ud.id
}
func (ud *userDomain) SetID(id string) {
	ud.id = id
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
