package main

import (
	"github.com/IshiniKiridena/VoteChainBackend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	//set up routes
	routes.HealthRouter(r)

	// Run the application on a specific port
	r.Run(":8080")
}
