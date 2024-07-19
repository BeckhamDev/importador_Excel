package controllers

import (
	"importador_Excel/internal/bcrypt"
	"importador_Excel/internal/database"
	"importador_Excel/internal/repository"
	"importador_Excel/internal/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var request LoginInput
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer db.Close()

	rep := repository.NewUserRepository(db)
	user, err := rep.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.ComparePassword(request.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := token.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}