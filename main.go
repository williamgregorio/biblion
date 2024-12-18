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

	for _, book := range bible.Books {
		if book.Book == "John" {
			for _, chapter := range book.Chapters {
				if chapter.Chapter == "3" {
					for _, verse := range chapter.Verses {
						if verse.Verse == "16" {
							fmt.Println("Book:", book.Book)
							fmt.Println("Chapter:", chapter.Chapter)
							fmt.Println("Verse:", verse.Verse)
							fmt.Println("Text:", verse.Text)
						}
					}
				}
			}
		}
	}
}
