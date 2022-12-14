package restAdapters

import (
	"github.com/gin-gonic/gin"
	baseServices "github.com/quangtran88/anifni-base/libs/services"
	"github.com/quangtran88/anifni-base/libs/utils"
	"github.com/quangtran88/anifni-gateway/adapters/services"
	"github.com/quangtran88/anifni-gateway/core/usecases"
)

func InitRoutes(r *gin.Engine) {
	env := baseUtils.GetEnvManager()

	userService := serviceAdapters.NewUserService(env)
	kafkaProducer := baseServices.NewKafkaProducer()

	authUseCase := usecases.NewAuthUseCase(userService, kafkaProducer)

	authHandler := NewAuthHandler(authUseCase)
	commonHandler := NewCommonHandler()

	r.GET("/ping", commonHandler.HandlePing)
	r.POST("/pre-register", authHandler.HandlePreRegister)
}
