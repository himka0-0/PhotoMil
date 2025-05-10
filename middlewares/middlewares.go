package middlewares

import (
	"MillyPhoto/db"
	"MillyPhoto/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		jwtSecret := os.Getenv("JWT_SECRET") // Загружаем секретный ключ из .env
		if jwtSecret == "" {
			log.Println("JWT_SECRET не найден в .env")
		}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		loginFromToken, ok := claims["login"].(string)
		if !ok {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		var user models.User
		err = db.DB.Where("login = ?", loginFromToken).First(&user).Error
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Set("Login", loginFromToken)
		c.Set("User", user)
		c.Next()
	}
}
