package request

type ShelterRequest struct {
	Name         string `json:"name" binding:"required" example:"Abrigo dos amigos"`
	City         string `json:"city" binding:"required" example:"Tallinn"`
	ZipCode      string `json:"zip_code" binding:"required" example:"14000000"`
	Address      string `json:"address" binding:"required" example:"моей улице"`
	Number       string `json:"number" binding:"required" example:"662"`
	Neighborhood string `json:"neighborhood" binding:"required" example:"Хааберсти"`
	Complement   string `json:"complement"`
}
