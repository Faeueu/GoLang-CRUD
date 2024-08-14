package controller

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Erro ao chamar rota!")
	c.JSON(err.Code, err)
}