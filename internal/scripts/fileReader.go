package scripts

import (
	"database/sql"
	"importador_Excel/internal/repository"
	"sync"

	"github.com/tealeg/xlsx"
)

func ReadFile(excelFileName string, ch chan<- []string) error {
	// Abre o arquivo Excel
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return err
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
					line = append(line, text)
				}
				ch <- line
			}
		}(sheet)
	}

	// Aguarda todas as goroutines produtoras terminarem
	wgProducers.Wait()
	close(ch)
	
	return nil
}

func ConsumeFileAndSaveData(ch <-chan []string, wg *sync.WaitGroup, db *sql.DB) error {
	defer wg.Done()
	var data [][]string // Array para armazenar os dados

	for val := range ch {
		sanitizedVal := Sanitize(val)
		data = append(data, sanitizedVal) // Salva os dados no array
	}

	//*****************TEMPORÁRIO*******************

	rep := repository.NewReconRepository(db)
	rep.SaveData(data)
	//*****************TEMPORÁRIO*******************

	return nil
}
