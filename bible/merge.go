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

func main() {

}
