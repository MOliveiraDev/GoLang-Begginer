package main

import "fmt"

func main() {

	idade := 21
	nome := "Pedro"
	pontuacao := 9.5

	fmt.Println("Olá, ")
	fmt.Println("mundo!")
	fmt.Println("Nova Linha")

	//PrintF (formatar string) %d (inteiro) %s (string) %v (valor generico) %f (numero flutuante) %t (booleano) %q (aspas em strings) %T (tipo do valor) - valor
	fmt.Printf("Minha idade é %d e meu nome é %s \n", idade, nome) 
	fmt.Printf("O tipo da variável idade é %T \n", idade)
	fmt.Printf("Minha pontuação é %f \n", pontuacao)
	fmt.Printf("O tipo de variável pontuação é %T \n", pontuacao)

	//Sprintf (formatar string e retornar a string formatada) %d (inteiro) %s (string) %v (valor generico) %f (numero flutuante) %t (booleano) %q (aspas em strings) %T (tipo do valor) - valor
	mensagem := fmt.Sprintf("Minha idade é %d e meu nome é %s", idade, nome)
	fmt.Println(mensagem)

}