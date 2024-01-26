package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


type Router struct {
	engine *gin.Engine
}

func main() {
	r := &Router{}
	r.engine = gin.Default()

	// Define routes
	// r.engine.GET("/users/:id/balance", api.GetUserBalanceHandler)

	// Run the server
	port := ":8081"
	fmt.Printf("Server running on %s\n", port)
	r.engine.Run(port)
}
