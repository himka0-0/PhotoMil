package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cabinet(c *gin.Context) {
	c.HTML(http.StatusOK, "Adminkab.html", gin.H{})
}
