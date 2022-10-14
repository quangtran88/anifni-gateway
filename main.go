package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/quangtran88/anifni-gateway/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}

	r := gin.Default()

	routes.InitCommonRoutes(r)
	routes.InitUserRoutes(r)

	err = r.Run(":5000")
	if err != nil {
		log.Println(err)
	}
}
