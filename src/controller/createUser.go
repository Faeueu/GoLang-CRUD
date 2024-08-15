package controller

import (
	"net/http"


	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/logs"
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/validation"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/request"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	logs.Info("Init createUser controller",
		zap.String("journey", "createUser"))

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logs.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)
		
		c.JSON(errRest.Code, errRest)
		return
	}
	
	logs.Info("User created successfully",
		zap.String("journey", "createUser"))	
	response := response.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}