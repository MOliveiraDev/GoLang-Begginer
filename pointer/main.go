package main

import "fmt"

func main() {

	nome1 := "Alberto"
	nome2 := "Richard"

	// O '&' captura o endereço de memória de 'name' e o armazena no ponteiro 'm'
	m1 := &nome1
	m2 := &nome2

	//Imprimir o valor hexadecimal da memória
	fmt.Println("Endereço de memória 'nome1':", m1)
	fmt.Println("Endereço de memória 'nome2':", m2)

	// O '*' desreferencia o ponteiro, indo até o endereço para ler o valor real
	fmt.Println("Valor no endereço 'nome1':", *m1)
	fmt.Println("Valor no endereço 'nome2':", *m2)

	// Atualizando o valor usando a função
	atualizaNome(m1)

	fmt.Println("Valor atualizado 'nome1':", nome1)
}

// Essa função recebe um ponteiro para string e atualiza o valor apontado para "Letícia"
func atualizaNome(nome *string) {
	*nome = "Letícia"
}