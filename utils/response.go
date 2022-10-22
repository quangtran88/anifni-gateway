package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReplyError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "failed",
		"error":  err.Error(),
	})
}

func ReplySuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func ReplyData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
