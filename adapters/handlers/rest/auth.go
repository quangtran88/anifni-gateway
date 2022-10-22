package restAdapters

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"github.com/quangtran88/anifni-gateway/utils"
)

type AuthHandler struct {
	authUC ports.AuthUseCase
}

func NewAuthHandler(authUC ports.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC}
}

func (h AuthHandler) HandlePreRegister(c *gin.Context) {
	var dto ports.PreRegisterUserInput
	err := c.BindJSON(&dto)
	if err != nil {
		utils.ReplyError(c, err)
		return
	}

	ok, err := h.authUC.PreRegisterUser(c.Copy(), dto)
	if err != nil || !ok {
		utils.ReplyError(c, fmt.Errorf("error while register user: %w", err))
		return
	}

	utils.ReplySuccess(c)
}
