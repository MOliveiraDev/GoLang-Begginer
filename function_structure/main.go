package main

import (
	"beginners/function_structure/models"
	"fmt"
)

func main() {
	cars := models.CarList()
	fmt.Println("Lista de Carros:")
	for _, car := range cars {
		// O \n dentro da string é usado para quebrar a linha
		fmt.Printf("Modelo: %s | Ano: %d\n", car.Model, car.Year)
	}
}
