package restHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}