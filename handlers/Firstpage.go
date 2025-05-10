package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FirstPage(c *gin.Context) {
	c.HTML(http.StatusOK, "FirstPage.html", gin.H{})
}
