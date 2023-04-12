package request

import "abrigos/source/domain/enumerations"

type UserRequest struct {
	Name     string             `json:"name" binding:"required" example:"Vinicius Rossado"`
	Password string             `json:"password" binding:"required" example:"Teste@100"`
	Username string             `json:"username" binding:"required" example:"vinirossado"`
	Role     enumerations.Roles `json:"role" binding:"required" example:"ADMIN"`
}
