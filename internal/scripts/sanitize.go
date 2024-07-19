package scripts

//Sanitize preenche as 3 ultimas colunas como vazias caso elas nao tenham sido reconhecidas pela funçao de ler o arquivo
func Sanitize(data []string) []string {

	if len(data) < 55 {
		emptydatas := 54 - len(data)
		for i := 0; i <= emptydatas; i++ {
			emptySpace := ""
			data = append(data, emptySpace)
		}
	}

	return data
}