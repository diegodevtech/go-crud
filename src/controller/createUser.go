package controller

import (
	"fmt"

	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		rest_error := rest_err.NewBadRequestError(fmt.Sprintf("There are some invalid fields. Error: %s", err.Error()))
		c.JSON(rest_error.Code, rest_error)
		return
	}

	fmt.Println(userRequest)
}