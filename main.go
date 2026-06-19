package main

import (
	"log"
	"net/http"

	"libary-management/internal/book"
	"libary-management/pkg/database"
	
)

func main() {
	db := database.InitDB()
	defer db.Close()

	repo := book.NewRepository(db)
	service := book.NewService(repo)
	handler := book.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", handler.GetBooks)
	mux.HandleFunc("GET /books/{id}", handler.GetBook)
	mux.HandleFunc("POST /books", handler.CreateBook)
	mux.HandleFunc("PUT /books/{id}", handler.UpdateBook)
	mux.HandleFunc("DELETE /books/{id}", handler.DeleteBook)

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", mux)
}
