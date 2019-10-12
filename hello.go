package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Book json structure
type book struct {
	IBN   string `json:"IBN"`
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}

// Books this is an array of books
var books []book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: books")
	json.NewEncoder(w).Encode(books)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", getBooks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	books = []book{
		book{IBN: "IBN123", Title: "Kubernetes for Dummies", Desc: "This book is intended to introduce Kubernetes to a novice."},
		book{IBN: "IBN456", Title: "Advanced Kubernetes", Desc: "This book delves deep into Kubernetes."},
	}

	handleRequests()
}
