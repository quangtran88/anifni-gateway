package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/quangtran88/anifni-gateway/core/services"
	rest "github.com/quangtran88/anifni-gateway/pkg/rest/handlers"
	"github.com/quangtran88/anifni-gateway/repositories"
)

func InitRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := rest.NewUserHandler(userService)

	r.GET("/ping", rest.HandlePing)
	r.GET("/user/ping", userHandler.HandlePingUser)
}
