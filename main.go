package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	var newbook Book
	err := json.NewDecoder(r.Body).Decode(&newbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newbook.ID = len(books) + 1

	books = append(books, newbook)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newbook)
}

func getbookbyID(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not Found", http.StatusNotFound)
}

func main() {
	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})
	books = append(books, Book{ID: 3, Title: "Book 3", Author: "Author 3"})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getBooks(w, r)
		case "POST":
			createBooks(w, r)
		case "DELETE":
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/book/{id}", getbookbyID)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
