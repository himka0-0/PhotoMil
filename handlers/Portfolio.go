package handlers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Portfolio(c *gin.Context) {
	var portfolios []models.Portfolio
	if err := db.DB.Preload("PortfolioPhoto").Find(&portfolios).Error; err != nil {
		log.Println("Ошибка при получении портфолио:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось загрузить портфолио",
		})
		return
	}
	type Album struct {
		Title  string
		Photos []string
	}

	var albums []Album
	for _, p := range portfolios {
		var photos []string
		for _, ph := range p.PortfolioPhoto {
			photos = append(photos, ph.Photo)
		}
		albums = append(albums, Album{
			Title:  p.Name,
			Photos: photos,
		})
	}

	c.HTML(http.StatusOK, "Portfolio.html", gin.H{
		"Albums": albums,
	})
}
