package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/a-h/templ"
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
	fmt.Fprintln(w, "Biblion")
}

func getVerse(bible Bible, bookName, chapter, verse string) (string, bool) {
	for _, book := range bible.Books {
		if book.Book == bookName {
			for _, chap := range book.Chapters {
				if chap.Chapter == chapter {
					for _, text := range chap.Verses {
						if text.Verse == verse {
							return text.Text, true
						}
					}
				}
			}
		}
	}
	return "", false
}

func main() {
	file, err := os.Open("./bible/Bible.json")
	if err != nil {
		fmt.Printf("error opening file: %s", err)
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("failed to read file: %s", err)
		return
	}

	var bible Bible

	if err := json.Unmarshal(byteValue, &bible); err != nil {
		fmt.Printf("failed to parse JSON: %s", err)
		return
	}

	bookName := "John"
	chapter := "3"
	verse := "16"

	text, found := getVerse(bible, bookName, chapter, verse)
	if found {
		fmt.Printf("%s %s:%s - %s\n", bookName, chapter, verse, text)
	} else {
		fmt.Printf("verse not found: %s %s:%s\n", bookName, chapter, verse)
	}

	component := hello("Biblion")
	http.Handle("/", templ.Handler(component))

	fmt.Println("listening on :7000")
	http.ListenAndServe(":7000", nil)
}
