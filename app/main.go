package main

import (
	"log"

	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/logs"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller"
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/routes"
	"github.com/Faeueu/GoLang-CRUD.git/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logs.Info("About to start user application")
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}