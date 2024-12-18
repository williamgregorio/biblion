package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Bible struct {
	Books []Book `json:"books"`
}

type Book struct {
	Book     string    `json:"book"`
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Chapter string  `json:"chapter"`
	Verses  []Verse `json:"verses"`
}

type Verse struct {
	Verse string `json:"verse"`
	Text  string `json:"text"`
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Biblion")
}

func handleAnotherRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "another router but with this string")
}

func main() {
	file, err := os.Open("./bible/Bible.json")
	if err != nil {
		fmt.Errorf("error opening file:", err)
		return
	}
	defer file.Close()

	byteVal, err := io.ReadAll(file)
	if err != nil {
		fmt.Errorf("filed to read file:", err)
		return
	}

	var bible Bible

	if err := json.Unmarshal(byteVal, &bible); err != nil {
		fmt.Errorf("failed to parse JSON:", err)
		return
	}

	fmt.Println(bible)
}
