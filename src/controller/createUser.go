package controller

import (
	"fmt"
	"log"

	//"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/validation"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/request"
	//"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	log.Println("Init CreateUsers controller")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)
		
		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Print(userRequest)
}