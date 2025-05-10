package handlers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminStudios(c *gin.Context) {
	var albom []models.Study
	if err := db.DB.Model(&models.Study{}).Find(&albom).Error; err != nil {
		log.Println("ошибка поиска альбьома в бд,AdminPortfolio", err)
	}
	c.HTML(http.StatusOK, "Adminstudios.html", gin.H{
		"alboms": albom,
	})
}
