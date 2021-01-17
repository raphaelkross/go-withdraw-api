package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelkross/go-withdraw-api/internal/api/controllers"
)

func main() {
	router := gin.Default()

	apiV1 := router.Group("/v1")

	// Add /withdraw/:amount route - accepting a int as arg.
	apiV1.GET("/withdraw/:amount", controllers.Withdraw)

	router.Run()
}
