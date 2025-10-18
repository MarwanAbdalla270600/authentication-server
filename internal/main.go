package main

import (
	"authentication-server/internal/auth"
	"authentication-server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	server := gin.Default()

	//database logic here
	db, err := sqlx.Open("sqlite", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = utils.CreateUserTable(db)
	if err != nil {
		panic(err)
	}

	controller := auth.NewController(nil)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})

	server.POST("/auth/register", controller.Register)
	server.POST("/auth/login", controller.Login)

	server.Run(":8080")

}
