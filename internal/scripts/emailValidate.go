package scripts

import "regexp"

func IsEmailValid(email string) bool {
	//Regex que valida se um email é válido
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}