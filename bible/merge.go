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
    return nil, err
  }
  var bookNames []string
  err = json.Unmarshal(data, &bookNames)
  return bookNames, err
}

func loadBookData(filename string) (*Book, error) {
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  var book Book
  err = json.Unmarshal(data, &book)
  return &book, err
}

func saveBibleData(filename string, bible *Bible) error {
  data, err := json.MarshalIndent(bible, "", "  ")
  if err != nil {
    return err
  }
  return ioutil.WriteFile(filename, data, 0644)
}

func main() {
  bookNames, err := loadBookNames("Books.json")
  if err != nil {
    log.Fatal("Failed to load Books.json: %v", err)
  }
}
