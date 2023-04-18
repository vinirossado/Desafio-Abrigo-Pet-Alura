package responses

type ShelterResponse struct {
	Name          string `json:"name"`
	City          string `json:"city"`
	ZipCode       string `json:"zip_code"`
	Address       string `json:"address"`
	Number        string `json:"number"`
	Neighborhood  string `json:"neighborhood"`
	Complement    string `json:"complement"`
	ResponsibleID uint   `json:"responsible_id"`
}
