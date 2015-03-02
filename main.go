package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// test to make sure the server is listening
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	carrier := router.Group("/carriers")

	{
		carrier.GET("/list", getCarriers)
		carrier.POST("create", createCarriers)
		carrier.DELETE("/delete", deleteCarriers)
		carrier.PUT("/update", updateCarriers)

	}

	circuit := router.Group("/circuits")
	{
		carrier.GET("/list", getCircuits)
		carrier.POST("/create", createCircuits)
		carrier.DELETE("/delete", deleteCircuits)
		carrier.POST("/update", updateCircuits)
	}
	// listen on port 3000
	r.Run(":3000")
}
