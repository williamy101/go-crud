package main

import (
	"go-asg4/config"
	"go-asg4/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()

	route.SetupRouter(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
