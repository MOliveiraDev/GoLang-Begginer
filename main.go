package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Imprime a mensagem "Olá mundo em GoLang!" no console
	fmt.Println("Olá mundo em GoLang!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o seu nome: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Nomes completos:")
	fmt.Println(name)

	// Definir idade, altura e peso
	fmt.Print("Digite a idade: ")
	var idade int
	fmt.Scanln(&idade)

	fmt.Print("Digite a altura (ex: 1,70 ou 1.70): ")
	alturaStr, _ := reader.ReadString('\n')
	alturaStr = strings.TrimSpace(alturaStr)
	alturaStr = strings.ReplaceAll(alturaStr, ",", ".")
	var altura float64
	fmt.Sscanf(alturaStr, "%f", &altura)

	fmt.Print("Digite o peso: ")
	var peso float64
	fmt.Scanln(&peso)

	// Imprime a idade
	fmt.Println("Idade:", idade)
	fmt.Printf("Altura: %.2f\n", altura)
	fmt.Println("Peso:", peso)
}
