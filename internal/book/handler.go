package book

import (
	"encoding/json"
	"errors"
	pkg "libary-management/pkg/response"
	"net/http"
	"strconv"
)

type ReqBook struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	ISBN       string `json:"isbn"`
	TotalCount int    `json:"total_count"`
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()

	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.WriteJSON(w, http.StatusOK, books)
}

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.service.GetBook(id)

	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteJSON(w, http.StatusOK, book)
}
func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req ReqBook

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book := Book{
		Title:          req.Title,
		Author:         req.Author,
		ISBN:           req.ISBN,
		TotalCount:     req.TotalCount,
		AvailableCount: req.TotalCount,
	}

	err := h.service.CreateBook(book)
	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, book)

}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	var req ReqBook

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book := Book{
		Title:      req.Title,
		Author:     req.Author,
		ISBN:       req.ISBN,
		TotalCount: req.TotalCount,
	}

	err = h.service.UpdateBook(id, book)

	if errors.Is(err, ErrorNotFound) {
		pkg.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	if errors.Is(err, ErrInvalidTotalCount) {
		pkg.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteJSON(w, http.StatusOK, book)
}
func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteBook(id)

	if errors.Is(err, ErrorNotFound) {
		pkg.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if errors.Is(err, ErrCannotDelete) {
		pkg.WriteError(w, http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.WriteJSON(w, http.StatusOK, id)
}
