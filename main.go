package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", getBooks)
	mux.HandleFunc("GET /books/{id}", getBook)
	mux.HandleFunc("POST /books", createBook)
	mux.HandleFunc("PUT /books/{id}", updateBook)
	mux.HandleFunc("DELETE /books/{id}", deleteBook)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]string{"status": "ok"})
	})

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}
