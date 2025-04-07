package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racosta/monorepo/projects/go_hello_world"
)

func RootHandler(c *gin.Context) {
	log.Println("Received request")
	c.JSON(http.StatusOK, gin.H{
		"message": go_hello_world.HelloWorld(),
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", RootHandler)

	// Bind to a port and pass our router in
	log.Println("Going to listen on port 8080")
	log.Fatal(r.Run(":8080"))
}
