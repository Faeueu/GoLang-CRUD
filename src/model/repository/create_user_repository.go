package repository

import (
	"context"
	"os"

	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/logs"
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"github.com/Faeueu/GoLang-CRUD.git/src/model"
)

const(
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logs.Info("Init createUser repository")

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue() 
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err  := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain,nil
}