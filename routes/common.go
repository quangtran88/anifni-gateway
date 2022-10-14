package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quangtran88/anifni-gateway/controllers"
)

func InitCommonRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.HandlePing)
}
