package main
import (
  "fmt"
  "net/http"
  )

  func handler( w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
  }

  func main() {
    fmt.Println("Starting server on port 3000")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
  }
