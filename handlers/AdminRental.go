package handlers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminRentals(c *gin.Context) {
	var albom []models.Rental
	if err := db.DB.Model(&models.Rental{}).Find(&albom).Error; err != nil {
		log.Println("ошибка поиска альбьома в бд,AdminPortfolio", err)
	}
	c.HTML(http.StatusOK, "Adminrental.html", gin.H{
		"clothes": albom,
	})
}
