package restHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"net/http"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler UserHandler) HandlePingUser(c *gin.Context) {
	result, err := handler.userService.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	}
}
