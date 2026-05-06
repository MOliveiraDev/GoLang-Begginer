package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/klauspost/cpuid/v2"
)

func main() {

	//Tempo de execução em milissegundos
	inicio := time.Now().UnixMilli()

	// O cpuid já inicializa os dados automaticamente
	fmt.Println("Nome do Processador:", cpuid.CPU.BrandName)
	fmt.Println("Número de Núcleos:", cpuid.CPU.PhysicalCores)
	fmt.Println("Threads por Núcleo:", cpuid.CPU.ThreadsPerCore)

	//Quantidade de núcleos trabalhando
	fmt.Println("Número de Processadores Lógicos:", runtime.NumCPU())

	//Definir o limite de nucleos a serem usados
	numeroDeNucleos := runtime.NumCPU() - 7 // Usando 1 núcleo a menos do que o total disponível
	runtime.GOMAXPROCS(numeroDeNucleos)

	var wgMain sync.WaitGroup
	wgMain.Add(1)

	//Função assíncrona para executar a tarefa
	go executarTarefa(numeroDeNucleos, &wgMain)

	wgMain.Wait()

	//Tempo total de execução do programa
	fim := time.Now().UnixMilli()
	fmt.Printf("Tempo Total do Programa: %d ms\n", fim-inicio)
}

func executarTarefa(numeroDeNucleos int, wgMain *sync.WaitGroup) {
	defer wgMain.Done()
	fmt.Printf("Número de Núcleos Definidos para Uso: %d\n", numeroDeNucleos)

	var wg sync.WaitGroup
	inicioGoroutines := time.Now()

	//Criar 10 goroutines para demonstrar o uso de múltiplos núcleos
	for i := 10; i > 0; i-- {
		wg.Add(1) // Adiciona 1 contador de espera para cada goroutine
		go func(i int) {
			defer wg.Done()        // Marca a goroutine como concluída ao final
			inicioGo := time.Now() // Tempo inicial desta goroutine

			time.Sleep(time.Duration(i) * 10 * time.Millisecond)
			if i == 10 {
				time.Sleep(100 * time.Microsecond)
			}

			duracao := time.Since(inicioGo) // Tempo que levou para rodar
			fmt.Printf("Goroutine %d terminou em %d ms\n", i, duracao.Milliseconds())
		}(i)
	}

	// Aguarda todas as goroutines terminarem (não precisamos mais de time.Sleep)
	wg.Wait()
	tempoTotalGoroutines := time.Since(inicioGoroutines)
	fmt.Printf("Tempo Total das Goroutines: %d ms\n", tempoTotalGoroutines.Milliseconds())
}
