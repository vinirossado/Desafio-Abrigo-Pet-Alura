package responses

import (
	"abrigos/source/domain/enumerations"
)

type UserResponse struct {
	ID       uint               `json:"id"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Role     enumerations.Roles `json:"role"`
}
