package main

import (
	"authentication-server/internal/auth"
	"authentication-server/internal/utils"

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

	// init singletons
	repo := auth.NewRepository(db)
	service := auth.NewService(repo)
	controller := auth.NewController(service)

	server.POST("/auth/register", controller.Register)
	server.POST("/auth/login", controller.Login)

	server.Run(":8080")

}
