package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostStruct struct {
	name    string
	message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	fmt.Println("Starting server on port 3000")
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":3000", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	input := json.NewDecoder(r.Body)
	var p PostStruct
	err := input.Decode(&p)
	if err != nil {
		panic(err)
	}

}
