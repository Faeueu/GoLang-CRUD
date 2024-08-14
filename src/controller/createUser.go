package controller

import (
	"fmt"

	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Campo incorreto, error=%s", err.Error()))
		
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}