package main

import (
  "fmt"
  "net/http"
)

type Book struct {
  Chapter string
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Biblion")
}

func main() {
  http.HandleFunc("/", handleRoot);

  fmt.Println("Server starting on port :7070")
  err := http.ListenAndServe(":7070", nil)
  if err != nil {
    fmt.Println("Server failed to start:", err)
  }
}
