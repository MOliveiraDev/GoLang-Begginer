package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var opcao int

	fmt.Println("Sistema de Gerenciamento de Produtos Gerais \n O que deseja fazer? \n 1 - Cadastrar produto \n 2 - Listar produtos \n 3 - Sair")


	fmt.Scanln(&opcao)
	switch opcao {
	case 1:
		cadastrarProduto(reader)
	case 2:
		listarProdutos()
	case 3:
		fmt.Println("Sair")
	default:
		fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
	}

}

func cadastrarProduto(reader *bufio.Reader) {

	fmt.Println("Digite o nome do produto:")

	nomeProduto, _ := reader.ReadString('\n')

	fmt.Println("Digite a quantidade do produto:")

	var quantidade int
	fmt.Scanln(&quantidade)

	fmt.Printf("Produto: %sQuantidade: %d\n", nomeProduto, quantidade)

	salvarProduto(nomeProduto, quantidade)

}

func listarProdutos() {
	startTime := time.Now()

	fmt.Println("Listando produtos cadastrados...")

	file, err := os.Open("produtos.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Printf("Tempo estimado: %v\n", time.Since(startTime))
}

func salvarProduto(nomeProduto string, quantidade int) {

	//registrando tempo de início da operação
	startTime := time.Now()

	//Caminho do arquivo onde os produtos serão salvos
	filePath := "produtos.txt"

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Produto: %sQuantidade: %d\n", nomeProduto, quantidade))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Produto cadastrado com sucesso!")

	//Tempo estimado da operação
	fmt.Printf("Tempo estimado para cadastro: %v\n", time.Since(startTime))

}