package main

import "fmt"

func main() {

	numero := ""
	//Solicitar ao usuário que digite um número decimal
	fmt.Println("Digite um número decimal")

	//Ler a entrada do usuário como string
	fmt.Scanln(&numero)

	//Colocando a string em um formato de número decimal usando Sscanf
	var numeroFloat float64

	// A função Sscanf lê a string 'numero' e tenta convertê-la para um número decimal, armazenando o resultado em 'numeroFloat'
	fmt.Sscanf(numero, "%f", &numeroFloat)

	fmt.Printf("O número digitado é: %.2f\n", numeroFloat)

}