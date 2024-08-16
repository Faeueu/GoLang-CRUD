package controller

import (
	"net/http"

	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/logs"
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/validation"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/request"
	"github.com/Faeueu/GoLang-CRUD.git/src/model"
	"github.com/Faeueu/GoLang-CRUD.git/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil{
		c.JSON(err.Code, err)
		return
	}
	
	logs.Info("User created successfully",
		zap.String("journey", "createUser"))	

	c.String(http.StatusOK, "")
	// response := response.UserResponse{
	// 	ID: "test",
	// 	Email: userRequest.Email,
	// 	Name: userRequest.Name,
	// 	Age: userRequest.Age,
	// }
}