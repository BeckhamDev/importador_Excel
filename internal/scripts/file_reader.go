package scripts

import (
	"fmt"
	"importador_Excel/internal/database"
	"importador_Excel/internal/repository"
	"log"
	"sync"

	"github.com/tealeg/xlsx"
)

func ReadFile(excelFileName string, ch chan<- []string) {
	// Abre o arquivo Excel
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo Excel: %s", err)
	}

	var wgProducers sync.WaitGroup

	// Itera sobre as planilhas
	for _, sheet := range xlFile.Sheets {
		wgProducers.Add(1)
		go func(sheet *xlsx.Sheet) {
			defer wgProducers.Done()
			// Itera sobre as linhas da planilha
			for rowIndex, row := range sheet.Rows {
				if rowIndex == 0 {
					continue
				}
	
				var line []string
				// Itera sobre as células da linha
				for _, cell := range row.Cells {
					
					text := cell.String() // Obtém o valor da célula como texto
					if text == ""{
						text = "teste"
					}

					line = append(line, text)
				}

				if len(line) < 55 {
					emptyLines := 54 - len(line)
					for i := 0; i <= emptyLines; i++ {
						text := "teste"

						line = append(line, text)
					}
				}

				ch <- line
			}
		}(sheet)
	}

	// Aguarda todas as goroutines produtoras terminarem
	wgProducers.Wait()
	close(ch)
}

func ConsumeFileAndSaveData(ch <-chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	var data [][]string // Array para armazenar os dados

	for val := range ch {
		data = append(data, val) // Salva os dados no array
	}

	//*****************TEMPORÁRIO*******************
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rep := repository.NewReconRepository(db)
	rep.SaveData(data)
	//*****************TEMPORÁRIO*******************
}
