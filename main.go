package main

import (
"fmt"
"log"
"github.com/gin-gonic/gin"
"encoding/json"
"github.com/boltdb/bolt"
)


type carrierJSON struct {
   CarrierName string 'json: "carrier_name"  binding: "required"'
   SupportEmail string 'json: "support_email"  binding: "required"'
   SupportNum string 'json: "support_num"  binding: "required"'
}

type circuitJSON struct {
  CircuitId string 'json: "circuit_id"  binding: "required"'
  CircuitLoc string 'json: "circuit_loc"  binding: "required"'
  CarrierBlock string 'json: "circuit_block"  binding: "required"'

}

func main() {
  // opden database
  db,err := bolt.Open("go-circuits.db", 0600, nil)
  if err != nil {
    log.Fatal(err)

  }
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
 func getCarriers(w h)
