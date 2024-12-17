package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/get", handleAnotherRoute)

	fmt.Println("Server starting on port :7070")
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
