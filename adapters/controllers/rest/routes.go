package restHandler

import (
	"github.com/gin-gonic/gin"
	restHandler "github.com/quangtran88/anifni-gateway/adapters/controllers/rest/handlers"
	"github.com/quangtran88/anifni-gateway/adapters/repositories"
	"github.com/quangtran88/anifni-gateway/core/services"
)

func InitRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := restHandler.NewUserHandler(userService)

	r.GET("/ping", restHandler.HandlePing)
	r.GET("/user/ping", userHandler.HandlePingUser)
	r.GET("/user/:id", userHandler.HandleGetUser)
}
