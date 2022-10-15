package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/quangtran88/anifni-gateway/adapters/controllers/rest"
	"github.com/quangtran88/anifni-gateway/constant"
	"github.com/quangtran88/anifni-gateway/utils"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}

	r := gin.Default()
	rest.InitRoutes(r)

	port := utils.GetEnvDefault(constant.GatewayPortEnvKey, "5000")
	log.Printf("Start gateway on http://localhost:%s", port)
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Cannot start gateway: %v", err)
	}
}
