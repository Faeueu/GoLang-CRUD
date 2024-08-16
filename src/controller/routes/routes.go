package routes

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface){

	

	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmai/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}