package routes

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmai/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}