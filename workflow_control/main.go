package main

import (
	"fmt"
	"os"
)

func main() {

	var idade int
	fmt.Println("Digite a sua idade:")

	//Verificando se a entrada é um número inteiro válido
	if _, err := fmt.Fscan(os.Stdin, &idade); err != nil {
		fmt.Println("Entrada inválida. Por favor, digite um número inteiro.")
		return
	}

	//Verificando se a idade é um valor positivo
	if idade < 0 {
		fmt.Println("Idade inválida. Por favor, insira um valor positivo.")
		return
	}

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
