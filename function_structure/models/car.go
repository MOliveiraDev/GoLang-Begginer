package models

import "time"

// Car representa uma entidade com os dados de um carro
type Car struct {
	Model           string
	Year            int
	ManufactureDate time.Time
}

func CarList() []Car {
	return []Car{
		{
			Model:           "Fusca",
			Year:            1980,
			ManufactureDate: time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Model:           "Civic",
			Year:            2020,
			ManufactureDate: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
}
