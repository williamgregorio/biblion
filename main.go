package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "html/template"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, err := loadPage(title)
  if err != nil {
    http.Error(w, "Page not found", http.StatusNotFound)
    return
  }
  fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  fmt.Fprintf(w, "<h1>Editing %s</h1>"+
    "<form action='/save/%s' method='POST'>"+
    "<textarea name='body'>%s</textarea><br>"+
    "<input type='submit' value='Save'/>"+
    "</form>", p.Title, p.Body)
}

func main()  {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", viewHandler)
  http.HandleFunc("/save/", viewHandler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
