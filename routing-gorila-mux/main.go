package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Printf("You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks).Methods("GET")
	bookrouter.HandleFunc("/{title}/", GetBook).Methods("GET")

	// Restrict the request handler to specific hostnames or subdomains.
	// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// Restrict the request handler to http/https.
	// r.HandleFunc("/secure", SecureHandler).Schemes("https")
	// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	http.ListenAndServe(":8081", r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateBook\n")
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ReadBook\n")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("UpdateBook\n")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DeleteBook\n")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("AllBooks\n")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetBooks\n")
}
