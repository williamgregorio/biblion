package main

import (
	"fmt"
	"net/http"
	"os"
)

type Library struct {
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
	fmt.Fprintf(w, "Another router but with this string")
}

func main() {
	file, err := os.Open("./bible/Bible.json")
	if err != nil {
		fmt.Errorf("error opening file:", err)
		return
	}

	fmt.Println(file)
	defer file.Close()
}
