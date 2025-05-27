package handlers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Rental(c *gin.Context) {
	var rantal []models.Rental
	if err := db.DB.Preload("RentalPhoto").Find(&rantal).Error; err != nil {
		log.Println("Ошибка при получении одежды:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось загрузить одежду",
		})
		return
	}

	type Album struct {
		Title  string
		Photos []string
	}

	var albums []Album
	for _, r := range rantal {
		var photos []string
		for _, ph := range r.RentalPhoto {
			photos = append(photos, ph.Photo)
		}
		albums = append(albums, Album{
			Title:  r.Name,
			Photos: photos,
		})
	}
	c.HTML(http.StatusOK, "rental.html", gin.H{
		"Albums": albums,
	})
}
