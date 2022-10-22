package restAdapters

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quangtran88/anifni-gateway/core/domain"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"net/http"
)

type AuthHandler struct {
	authUC ports.AuthUseCase
}

func NewAuthHandler(authUC ports.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC}
}

func (handler AuthHandler) HandleRegister(c *gin.Context) {
	var dto domain.RegisterUserInput
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ok, err := handler.authUC.RegisterUser(c.Copy(), dto)
	if err != nil || !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error while register user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}
