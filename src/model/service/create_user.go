package service

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/logs"
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"github.com/Faeueu/GoLang-CRUD.git/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logs.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}