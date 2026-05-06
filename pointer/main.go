package main

import "fmt"

func main() {

	nome := "Alberto"

	// O '&' captura o endereço de memória de 'name' e o armazena no ponteiro 'm'
	m := &nome

	//Imprimir o valor hexadecimal da memória
	fmt.Println("Endereço de memória 'nome':", m)

	// O '*' desreferencia o ponteiro, indo até o endereço para ler o valor real
	fmt.Println("Valor no endereço:", *m)
}