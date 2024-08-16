package model

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/logs"
	"github.com/Faeueu/GoLang-CRUD.git/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logs.Info("Init create user model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}