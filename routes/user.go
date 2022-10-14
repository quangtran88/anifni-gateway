package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quangtran88/anifni-gateway/controllers"
)

func InitUserRoutes(r *gin.Engine) {
	r.GET("/user/ping", controllers.HandlePingUser)
}
