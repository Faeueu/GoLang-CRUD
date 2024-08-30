package repository

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"github.com/Faeueu/GoLang-CRUD.git/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository (
	database *mongo.Database,
) UserRepository {
	return &userRepository {
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomais model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

}