package controllers

import (
	"importador_Excel/internal/bcrypt"
	"importador_Excel/internal/database"
	"importador_Excel/internal/models"
	"importador_Excel/internal/repository"
	"importador_Excel/internal/scripts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//Verifica se o email está no formato adequado
	if !scripts.IsEmailValid(request.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email inválido!"})
		return
	}


	//Inicia instancia de conexao com o banco
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	rep := repository.NewUserRepository(db)
	user, err := rep.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//Verifica se o email já possui cadastro
	if (models.User{} != user) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Este email já possui um cadastro"})
		return
	}

	//Criptgrafa a senha para que possa ser salva
	hashedPassword, err := bcrypt.Hash(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	request.Password = string(hashedPassword)

	//Cria o usuário no banco de dados
	user, err = rep.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}