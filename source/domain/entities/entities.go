package entities

func RetriveAll() []interface{} {
	return []interface{}{
		&User{},
		&Shelter{},
		&Pet{},
		&Adoption{},
	}
}
