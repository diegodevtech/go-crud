package model

import (
	// "crypto/md5"
	// "encoding/hex"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomain) EncryptPassword() {
	// hash := md5.New()
	// defer hash.Reset()
	// hash.Write([]byte(ud.password))
	// ud.password = hex.EncodeToString(hash.Sum(nil))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Fail attempting to encrypt password", err)
	}
	ud.password = string(hashedPassword)
}