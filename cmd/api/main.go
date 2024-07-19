package main

import (
	"fmt"
	"importador_Excel/config"
	"importador_Excel/internal/router"
)

func main() {

	fmt.Println("Rodando")
	config.Load()
	router.Initialize()

}