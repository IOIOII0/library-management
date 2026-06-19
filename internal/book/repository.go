package book

import (
	"database/sql"
	"errors"
	"fmt"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Book, error) {
	rows, err := r.db.Query(
		"SELECT id, title, author, isbn, total_count, available_count FROM books")

	if err != nil {
		return nil, fmt.Errorf("find all book : %w", err)
	}
	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		var b Book
		err := rows.Scan(
			&b.ID, &b.Title, &b.Author, &b.ISBN, &b.TotalCount, &b.AvailableCount)
		if err != nil {
			return nil, fmt.Errorf("scan books rows %w", err)
		}
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate book rows: %w", err)
	}
	return books, nil
}

func (r *repository) FindByID(id int) (Book, error) {

	var b Book
	err := r.db.QueryRow(
		"SELECT id, title, author, isbn, total_count, available_count FROM books WHERE id = $1",
		id).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.TotalCount, &b.AvailableCount)

	if errors.Is(err, sql.ErrNoRows) {
		return Book{}, ErrorNotFound
	}

	if err != nil {
		return Book{}, fmt.Errorf("find book By id : %w", err)
	}

	return b, nil
}

func (r *repository) Create(b Book) error {

	_, err := r.db.Exec(
		"INSERT INTO books (title,author,isbn,total_count, available_count) VALUES ($1, $2, $3, $4, $5)",
		b.Title, b.Author, b.ISBN, b.TotalCount, b.AvailableCount)

	if err != nil {
		return fmt.Errorf("Insert book : %w", err)
	}
	return nil
}

func (r *repository) Update(id int, b Book) error {

	_, err := r.db.Exec(
		"UPDATE books SET title = $1, author = $2, isbn = $3, total_count = $4, available_count =$5 WHERE id = $6",
		b.Title, b.Author, b.ISBN, b.TotalCount, b.AvailableCount, id)

	if err != nil {
		return fmt.Errorf("update book %d : %w", id, err)
	}

	return nil
}
func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Delete book from id : %d id : %w", id, err)
	}
	return nil
}
