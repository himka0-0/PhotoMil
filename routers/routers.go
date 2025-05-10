package routers

import (
	"MillyPhoto/controllers"
	"MillyPhoto/handlers"
	"MillyPhoto/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {
	r.GET("/", handlers.FirstPage)
	r.GET("/portfolio", handlers.Portfolio)
	r.GET("/studios", handlers.Study)
	r.GET("/rental", handlers.Rental)
	r.GET("/booking", handlers.Booking)
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.Login)
	admin := r.Group("/user")
	admin.Use(middlewares.AuthMiddleware())
	{
		admin.GET("/cabinet", controllers.Cabinet)

		admin.GET("/portfolios", handlers.AdminPortfolio)
		admin.POST("/portfolios", controllers.UploadAlbum)
		admin.DELETE("/portfolios", controllers.DeleteAlbom)

		admin.GET("/studioss", handlers.AdminStudios)
		admin.POST("/studioss", controllers.UploadStudios)
		admin.DELETE("/studioss", controllers.DeleteStudios)

		admin.GET("/rentals", handlers.AdminRentals)
		admin.POST("/rentals", controllers.UploadRental)
		admin.DELETE("/rentals", controllers.DeleteRental)
	}
}
