package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
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
type server struct {
  db *bolt.DB
}

func main() {
  db,err := bolt.Open("test.db", 0600, nil)
  if err != nil {
    log.Fatal(err)
  }
  
  s := server{
    db : db,
  }
  
  s.initialize()
  router := gin.Default()
  router.GET("/ping",func(c *gin.Context){
    c.String(200, "pong")
    })
  router.GET("/circuits", s.getCircuits)
}

func (s *server) getCircuits(c *gin.Context) {
  s.db.Update(func(tx *bolt.Tx) error {
    b:= tx.Bucket([]byte('circuits'))
    b.ForEach(func(key, value []byte) error {
      log.Println(string(key), string(value))
      
    })
    
  })
}

func (s *server) postCircuits(c *gin.Context){
  var cirs circuitJSON
  err :=  c.Bind(&cirs)
  if err != nil {
    // inbound json was invalid
    log.Println("Error parsing JSON:", err)
  }
    s.db.Update(func(tx *bolt.Tx) error {
      b:= tx.Bucket([]byte('circuits'))
      b.Put([]byte(cirs.CircuitID), []byte(cirs.CircuitLoc), []byte(cirs.CarrierBlock))
      return nil
      
    })
    fmt.Println(cirs)
}
// Do anything you need to setup before the server actually runs
func (s *server) initialize(){
  s.db.Update(func(tx *bolt.Tx) error {
    // Make sure you create the buckets you need,
    // if you try to access a bucket which doesn't exist
    // you will get a panic
    tx.CreateBucket([]byte(`circuits`))
    return nil
  })  
}
