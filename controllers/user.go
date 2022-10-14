package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func HandlePingUser(c *gin.Context) {
	conn, err := grpc.Dial("localhost:6000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := userGRPC.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(c.Copy(), time.Second)
	defer cancel()

	r, err := client.Ping(ctx, &userGRPC.PingMessage{})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": r.GetMessage(),
	})
}
