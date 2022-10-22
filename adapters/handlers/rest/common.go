package restAdapters

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonHandler struct {
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{}
}

func (handler CommonHandler) HandlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
