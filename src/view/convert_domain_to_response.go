package view

import (
	"github.com/Faeueu/GoLang-CRUD.git/src/controller/model/response"
	"github.com/Faeueu/GoLang-CRUD.git/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}