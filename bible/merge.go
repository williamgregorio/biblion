package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "path/filepath"
  "strings"
)

type Verse struct {
  Verse string `json:"verse"`
  Text string `json:"text"`
}

type Chapter struct {
  Chapter string `json:"chapter"`
  Verses []Verse `json:"verses"`
}

type Book struct {
  Book string `json:"book"`
  Chapters []Chapter `json:"chapters"`
}

type Bible struct {
  Books []Book `json:"books"`
}

func loadBookNames(filename string) ([]string, error) {
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    // use a proper message later
    return nil, err
  }
  var bookNames []string
  err = json.Unmarshal(data, &bookNames)
  return bookNames, err
}

func main() {
  bookNames, err := loadBookNames("Books.json")
  if err != nil {
    log.Fatal("Failed to load Books.json: %v", err)
  }
}
