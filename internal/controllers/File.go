package controllers

import (
	"importador_Excel/internal/database"
	"importador_Excel/internal/repository"
	"importador_Excel/internal/scripts"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func UploadFileAndSaveData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Especifica o caminho onde o arquivo será salvo
	filePath := "../../arquivo_teste.xlsx"

	// Salva o arquivo no disco
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Cria instancia de conexao com o banco
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	// Cria canal de dados
	ch := make(chan []string)

	var wg sync.WaitGroup

	// Número de consumidores
	numConsumers := 4

	// Adiciona o número de consumidores ao grupo de espera
	wg.Add(numConsumers)

	// Inicia as goroutines consumidoras
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := scripts.ConsumeFileAndSaveData(ch, &wg, db); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}()
	}

	// Lê o arquivo Excel e envia os dados para o canal
	scripts.ReadFile(filePath, ch)

	// Espera todas as goroutines consumidoras terminarem
	wg.Wait()

	//Apaga arquivo após salvar os dados
	err = os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dados salvos com sucesso!"})
}

func GetRowByCustomerName(c *gin.Context) {
	customerName := c.Param("customerName")
	if customerName == "" || customerName == " " {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Identificador inválido"})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	rep := repository.NewReconRepository(db)
	rows, err := rep.GetByCustomerName(customerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Customers": rows})
}


func GetRowByMeterCategory(c *gin.Context) {
	meterCategory := c.Param("meterCategory")
	if meterCategory == "" || meterCategory == " " {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Identificador inválido"})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	rep := repository.NewReconRepository(db)
	rows, err := rep.GetByMeterCategory(meterCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Customers": rows})
}

func GetRowByMeterRegion(c *gin.Context) {
	meterRegion := c.Param("meterRegion")
	if meterRegion == "" || meterRegion == " " {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Identificador inválido"})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	rep := repository.NewReconRepository(db)
	rows, err := rep.GetByMeterRegion(meterRegion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Customers": rows})
}

func GetRowByResourceGroup(c *gin.Context) {
	resourceGroup := c.Param("resourceGroup")
	if resourceGroup == "" || resourceGroup == " " {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Identificador inválido"})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	rep := repository.NewReconRepository(db)
	rows, err := rep.GetByResourceGroup(resourceGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Customers": rows})
}