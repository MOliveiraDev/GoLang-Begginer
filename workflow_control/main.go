package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Digite a sua idade:")
	fmt.Scanln(&idade)
	
	// Estrutura de controle if-else para verificar a idade
	if idade >= 18 {
		println("Você é maior de idade.")
	} else if idade >= 13 {
		println("Você é adolescente.")
	} else {
		println("Você é criança.")
	}

	// Usando switch para verificar a idade
	switch {
	case idade >= 18:
		println("Você é maior de idade.")
	case idade >= 13:
		println("Você é adolescente.")
	default:
		println("Você é criança.")
	}
}
