package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() err {
  filename := p.title + ".txt"
  err := os.WriteFile(filename, p.Body, 0600)
  if err != nil {
    fmt.Errorf("failed to save page: %w",err)
  }
  return nil
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := os.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Biblion %s", r.URL.Path)
}

func main()  {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
