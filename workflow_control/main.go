package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Digite a sua idade:")
	fmt.Scanln(&idade)
	
	// Estrutura de controle if-else para verificar a idade
	if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else if idade >= 13 {
		fmt.Println("Você é adolescente.")
	} else {
		fmt.Println("Você é criança.")
	}

	// Usando switch para verificar a idade
	switch {
	case idade >= 18:
		fmt.Println("Você é maior de idade.")
	case idade >= 13:
		fmt.Println("Você é adolescente.")
	case idade < 13:
		fmt.Println("Você é criança.")

	default:
		fmt.Println("Valor inválido.")
	}
}
