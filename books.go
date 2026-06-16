package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type Book struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	ISBN           string `json:"isbn"`
	TotalCount     int    `json:"total_count"`
	AvailableCount int    `json:"available_count"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author, isbn, total_count, available_count FROM books")

	if err != nil {
		writeError(w, 500, "query failed")
		return
	}
	defer rows.Close()

	var books = []Book{}
	for rows.Next() {
		var b Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.TotalCount, &b.AvailableCount)
		books = append(books, b)
	}
	writeJSON(w, http.StatusOK, books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	var b Book
	err = db.QueryRow("SELECT id, title, author, isbn, total_count, available_count FROM books WHERE id = $1", id).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.TotalCount, &b.AvailableCount)

	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query failed")
		return
	}
	writeJSON(w, http.StatusOK, b)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title      string `json:"title"`
		Author     string `json:"author"`
		ISBN       string `json:"isbn"`
		TotalCount int    `json:"total_count"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	_, err := db.Exec("INSERT INTO books (title, author, isbn, total_count, available_count) VALUES ($1, $2, $3, $4, $4)",
		body.Title, body.Author, body.ISBN, body.TotalCount)

	if err != nil {
		writeError(w, 500, "insert failed")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "book create"})

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title      string `json:"title"`
		Author     string `json:"author"`
		ISBN       string `json:"isbn"`
		TotalCount int    `json:"total_count"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	_, err = db.Exec("UPDATE books SET title = $1, author = $2, isbn = $3, total_count = $4, available_count = $4 WHERE id = $5",
		body.Title, body.Author, body.ISBN, body.TotalCount, id)

	if err != nil {
		writeError(w, 500, "insert failed")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "updated"})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	_, err = db.Exec("DELETE FROM books WHERE id = $1", id)

	if err != nil {
		writeError(w, 500, "delete failed")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "delete Success"})
}
