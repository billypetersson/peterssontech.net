package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define your routes and handlers here

	r.Run(":8080")
}
