package book

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(b Book) error
	Update(id int, b Book) error
	Delete(id int) error
}

type Service interface {
	GetBooks() ([]Book, error)
	GetBook(id int) (Book, error)
	CreateBook(b Book) error
	UpdateBook(id int, b Book) error
	DeleteBook(id int) error
}
