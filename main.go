package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

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

	fmt.Println("Starting server on port 3000")
	http.HandleFunc("/ping", s.getPing)
	http.HandleFunc("/change", s.postChange)
	http.ListenAndServe(":3000", nil)
}

type server struct {
	db *bolt.DB
}

// This is for a ping request.
func (s *server) getPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// This is to handle the 'change' route. Will determine if POST or GET
func (s *server) change(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		s.postChange(w, r)
	} else if r.Method == "GET" {
		s.getChange(w, r)
	}
}

// This is for getting a change request
func (s *server) getChange(w http.ResponseWriter, r *http.Request) {
	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`changes`))

		b.ForEach(func(key, value []byte) error {
			log.Println(string(key), string(value))

			return nil
		})

		return nil
	})
}

// This is for posting a change request
func (s *server) postChange(w http.ResponseWriter, r *http.Request) {
	var c change

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		// We are doing a console log, rather than panicing. This usually means the
		// inbound json was invalid
		log.Println("Error parsing JSON:", err)
		return
	}
	//check to make sure the change struct is not nil
	if c.Name == " " || c.Message == " " {
		w.Write([]byte("Invalid input"))
	}
	s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(`changes`))

		// When you put a key/value into bolt, it needs to be
		// in the form of a byteslice ([]byte).
		// Usually here, you would marshal your struct into
		// json, but we are going to keep it simple for now.
		b.Put([]byte(c.Name), []byte(c.Message))
		return nil
	})
	fmt.Println(c)
}

// Do anything you need to setup before the server actually runs
func (s *server) initialize() {
	s.db.Update(func(tx *bolt.Tx) error {
		// Make sure you create the buckets you need,
		// if you try to access a bucket which doesn't exist
		// you will get a panic
		tx.CreateBucket([]byte(`changes`))
		return nil
	})
}

// When you progress, this will be called when you catch a SIGINT.
// For now, it will just sit here
func (s *server) close() {
	s.db.Close()
}

// Generally, I try to keep my structs private unless I know I will
// be exporting them.
type change struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
