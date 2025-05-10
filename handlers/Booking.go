package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Booking(c *gin.Context) {
	c.HTML(http.StatusOK, "booking.html", gin.H{})
}
