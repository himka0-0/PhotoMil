package handlers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminPortfolio(c *gin.Context) {
	var albom []models.Portfolio
	if err := db.DB.Model(&models.Portfolio{}).Find(&albom).Error; err != nil {
		log.Println("ошибка поиска альбьома в бд,AdminPortfolio", err)
	}
	c.HTML(http.StatusOK, "Adminalbom.html", gin.H{
		"alboms": albom,
	})
}
