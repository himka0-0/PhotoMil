package controllers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"MillyPhoto/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Adminauth.html", gin.H{})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("ошибка парсинга при аутентификации")
	}
	var user models.User
	if er := db.DB.Where("login=?", input.Login).First(&user).Error; er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	token := utils.GenerateJwt(user.Login)
	c.SetCookie("token", token, 3600*24*7, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
