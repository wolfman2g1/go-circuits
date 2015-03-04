package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
)

type carrier struct {
	CarrierName  string `json: "carrier_name"  binding: "required"`
	SupportEmail string `json: "support_email"  binding: "required"`
	SupportNum   string `json: "support_num"  binding: "required"`
}

type circuit struct {
	CircuitId    string `json: "circuit_id"  binding: "required"`
	CircuitLoc   string `json: "circuit_loc"  binding: "required"`
	CarrierBlock string `json: "circuit_block"  binding: "required"`
}
type server struct {
	db *bolt.DB
}

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := server{
		db: db,
	}
	defer s.close()

	s.initialize()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// circuit end point
	router.GET("/circuit", s.getCircuits)
	router.POST("/circuit", s.postCircuits)

	// carrier end point
	router.GET("/carrier", s.getCarriers)
	router.POST("/carrier", s.postCarriers)
	// listen on this port
	router.Run(":3000")
}

func (s *server) getCircuits(c *gin.Context) {
	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`circuits`))
		b.ForEach(func(key, value []byte) error {
			log.Println(string(key), string(value))
			return

		})
		return

	})

}

func (s *server) postCircuits(c *gin.Context) {
	var cirs circuit
	if !c.Bind(&cirs) {

		// inbound json was invalid
		c.String(200, "Error parsing JSON:")
		return

	}
	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`circuits`))
		b.Put([]byte(cirs.CircuitId), []byte(cirs.CircuitLoc), []byte(cirs.CarrierBlock))
		return

	})
	fmt.Println(cirs)
}

// get a list of carriers
func (s *server) getCarriers(c *gin.Context) {
	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`carriers`))
		b.ForEach(func(key, value []byte) error {
			log.Println(string(key), string(value))
			return
		})
		return
	})
}

// create carriers
func (s *server) postCarriers(c *gin.Context) {
	var car carrier
	if !c.Bind(&car) {
		// inbound json was invalid
		c.String(200, "Error parsing JSON:")
		return

	}

	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`circuits`))
		b.Put([]byte(car.CarrierName), []byte(car.SupportEmail), []byte(car.SupportNum))
		return
	})
	fmt.Println(car)
}

// Do anything you need to setup before the server actually runs
func (s *server) initialize() {
	s.db.Update(func(tx *bolt.Tx) error {
		// Make sure you create the buckets you need,
		// if you try to access a bucket which doesn't exist
		// you will get a panic
		tx.CreateBucket([]byte(`circuits`))
		tx.CreateBucketIfNotExists([]byte(`carriers`))
		return
	})
}

func (s *server) close() {
	s.db.Close()
}
