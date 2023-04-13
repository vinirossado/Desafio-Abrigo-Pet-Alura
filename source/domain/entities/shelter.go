package entities

import "gorm.io/gorm"

type Shelter struct {
	gorm.Model

	Name   string ``
	Cidade string ``
}
