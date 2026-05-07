package main

import (
	"fmt"
	"net/http"
)

func main () {

	// Inicia o servidor HTTP na porta 8080
	fmt.Println("Servidor rodando em http://localhost:8080")

	// Inicia o servidor HTTP e escuta por requisições
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo em GoLang!\n"))
}