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
  return os.Writefile(filename, p.Body, 0600)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Biblion %s", r.URL.Path)
}

func main()  {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
