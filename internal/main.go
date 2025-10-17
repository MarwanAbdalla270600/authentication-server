package main

import (
	"authentication-server/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	controller := auth.NewController(nil)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})

	server.POST("/auth/register", controller.Register)
	server.POST("/auth/login", controller.Login)

	server.Run(":8080")

}
