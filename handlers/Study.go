package handlers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Study(c *gin.Context) {
	var study []models.Study
	if err := db.DB.Preload("StudyPhoto").Find(&study).Error; err != nil {
		log.Println("Ошибка при получении студий:", err)
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

	for _, st := range study {
		var photos []string
		for _, ph := range st.StudyPhoto {
			photos = append(photos, ph.Photo)
		}
		albums = append(albums, Album{
			Title:  st.Name,
			Photos: photos,
		})
	}
	c.HTML(http.StatusOK, "Study.html", gin.H{
		"Albums": albums,
	})
}
