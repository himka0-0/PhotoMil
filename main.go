package main

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"MillyPhoto/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	db.InitDB()
	db.DB.AutoMigrate(
		&models.User{Login: os.Getenv("LOGIN"), Password: os.Getenv("PASSWORD")},
		&models.Portfolio{}, &models.PortfolioPhoto{},
		&models.Study{}, &models.StudyPhoto{},
		&models.Rental{}, &models.RentalPhoto{})
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/uploads", "./uploads")
	r.Static("/study", "./study")
	r.Static("/rental", "./rental")

	routers.SetupRouters(r)
	r.Run(":8082")
}
