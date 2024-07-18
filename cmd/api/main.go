package main

import (
	"importador_Excel/internal/scripts"
	"sync"
)

func main() {
	excelFileName := "./arquivo_teste.xlsx"

	ch := make(chan []string) // Canal com buffer

	var wg sync.WaitGroup

	// Número de consumidores
	numConsumers := 4

	// Adiciona o número de consumidores ao grupo de espera
	wg.Add(numConsumers)

	// Inicia as goroutines consumidoras
	for i := 0; i < numConsumers; i++ {
		go scripts.ConsumeFileAndSaveData(ch, &wg)
	}

	// Lê o arquivo Excel e envia os dados para o canal
	scripts.ReadFile(excelFileName, ch)

	// Espera todas as goroutines consumidoras terminarem
	wg.Wait()

}