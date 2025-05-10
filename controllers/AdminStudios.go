package controllers

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func UploadStudios(c *gin.Context) {
	name := c.PostForm("name")
	form, _ := c.MultipartForm()
	files := form.File["photos"]

	if err := db.DB.Model(&models.Study{}).Create(&models.Study{Name: name}).Error; err != nil {
		log.Println("ошибка сохранения альбома", err)
		c.JSON(http.StatusConflict, gin.H{"error": "ошибка сохранения данных"})
		return
	}
	var albom models.Study
	err := db.DB.Where("name = ?", name).First(&albom).Error
	if err != nil {
		log.Println("ошибка поиске(сохранении) альбома", err)
		c.JSON(http.StatusConflict, gin.H{"error": "ошибка сохранения данных"})
		return
	}
	namber := 1
	for _, file := range files {
		ext := strings.TrimPrefix(filepath.Ext(file.Filename), ".")
		file.Filename = fmt.Sprintf("%s-%s.%s", name, strconv.Itoa(namber), ext)
		path := fmt.Sprintf("study/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(500, gin.H{"error": "Ошибка при сохранении файла"})
			return
		}

		db.DB.Create(&models.StudyPhoto{
			StudyID: albom.ID,
			Photo:   "/" + path,
		})
		namber++
	}

	c.JSON(200, gin.H{"success": true})
}

func DeleteStudios(c *gin.Context) {
	var payload struct {
		ID uint `json:"id"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	var photos []models.StudyPhoto
	if err := db.DB.Where("study_id = ?", payload.ID).Find(&photos).Error; err != nil {
		log.Println("Ошибка при получении фото:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении фото"})
		return
	}

	for _, photo := range photos {
		if err := os.Remove("." + photo.Photo); err != nil {
			log.Printf("Не удалось удалить файл %s: %v\n", photo.Photo, err)
		}
	}

	if err := db.DB.Where("study_id = ?", payload.ID).Delete(&models.StudyPhoto{}).Error; err != nil {
		log.Println("Ошибка при удалении фото из БД:", err)
	}

	if err := db.DB.Delete(&models.Study{}, payload.ID).Error; err != nil {
		log.Println("Ошибка при удалении альбома:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления альбома"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
