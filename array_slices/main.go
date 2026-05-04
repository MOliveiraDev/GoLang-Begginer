package main

import "fmt"

func main() {

	idades := [3]int{21, 22, 23} // Array de inteiros com 3 elementos
	fmt.Println("Array de idades:", idades, "\n Quantidade:", len(idades))

	nomes := [3]string{"Carle", "Bob", "Charlie"} // Array de strings com 3 elementos
	fmt.Println("Array de nomes:", nomes, "\n Quantidade:", len(nomes))

	// Acessando elementos do array
	fmt.Println("Primeira idade:", idades[0]) // Acessa o primeiro elemento do array de idades
	fmt.Println("Segunda idade:", idades[1])  // Acessa o segundo elemento do array de idades
	fmt.Println("Primeiro nome:", nomes[0])   // Acessa o primeiro elemento do array de nomes
	fmt.Println("Segundo nome:", nomes[1])    // Acessa o segundo elemento do array de nomes

	// Slice em float64 para armazenar pontuações
	pontuacoes := []float64{85.5, 90.0, 78.3} // Slice de float64 com 3 elementos
	fmt.Println("Slice de pontuações:", pontuacoes, "\n Quantidade:", len(pontuacoes))

}