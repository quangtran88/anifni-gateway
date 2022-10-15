package rest

import (
	"github.com/gin-gonic/gin"
	rest "github.com/quangtran88/anifni-gateway/adapters/controllers/rest/handlers"
	"github.com/quangtran88/anifni-gateway/adapters/repositories"
	"github.com/quangtran88/anifni-gateway/core/services"
)

func InitRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := rest.NewUserHandler(userService)

	r.GET("/ping", rest.HandlePing)
	r.GET("/user/ping", userHandler.HandlePingUser)
}
