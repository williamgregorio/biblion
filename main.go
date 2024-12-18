package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/williamgregorio/biblion/views/layout"
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

// helper - load bible
var bible Bible

func loadBible() error {
	file, err := os.Open("./bible/Bible.json")
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}

	if err := json.Unmarshal(byteValue, &bible); err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}

	return nil
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	root := layout.Base()
	if err := root.Render(ctx, w); err != nil {
		http.Error(w, "failed to render base template", http.StatusInternalServerError)
		fmt.Printf("error rendering template: %s\n", err)
	}
}

func handleGetVerse(w http.ResponseWriter, req *http.Request) {
	// qry params parse
	bookName := req.URL.Query().Get("book")
	chapter := req.URL.Query().Get("chapter")
	verse := req.URL.Query().Get("verse")

	if bookName == "" || chapter == "" || verse == "" {
		http.Error(w, "missing, required query params: book, chapter, verse", http.StatusBadRequest)
		return
	}

	// fetch verse
	text, found := getVerse(bible, bookName, chapter, verse)
	if !found {
		http.Error(w, "verse not found", http.StatusNotFound)
		return
	}

	// json return
	response := map[string]string{
		"bookName": bookName,
		"chapter":  chapter,
		"verse":    verse,
		"text":     text,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	if err := loadBible(); err != nil {
		fmt.Printf("failed to load bible: %s\n", err)
		return
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/verse", handleGetVerse)

	fmt.Println("listening on http://localhost:7000")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		fmt.Printf("failed to start web server: %s\n", err)
	}

}
