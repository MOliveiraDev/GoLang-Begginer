package main

import (
	"fmt"
	"os"
	"time"
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
	// switch {
	// case idade >= 18:
	// 	fmt.Println("Você é maior de idade.")
	// case idade >= 13:
	// 	fmt.Println("Você é adolescente.")
	// case idade < 13:
	// 	fmt.Println("Você é criança.")

	// default:
	// 	fmt.Println("Valor inválido.")
	// }

	// Usando if para identificar o ano de nascimento
	anoAtual := time.Now().Year()
	anoNascimento := anoAtual - idade
	if anoNascimento >= 0 {
		fmt.Printf("Você nasceu em %d\n", anoNascimento)
	} else {
		fmt.Println("Ano de nascimento inválido.")
	}


	// Usando for sobre Slice

	series := []string{"Breaking Bad", "Game of Thrones", "Stranger Things"}
	fmt.Println("Minhas séries favoritas:")
	for i, serie := range series {
		fmt.Printf("%d. %s\n", i+1, serie)
	}

	// Usando for para contar de 1 a 10
	fmt.Println("Contando de 1 a 10:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
