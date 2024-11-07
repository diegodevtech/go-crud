package routes

import (
	"github.com/diegodevtech/go-crud/src/controller"
	"github.com/gin-gonic/gin"
)



func initRoutes(r *gin.RouterGroup) {
	r.GET("/getUsersById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}