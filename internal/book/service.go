package book

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (r *service) GetBooks() ([]Book, error) {
	books, err := r.repo.FindAll()

	if err != nil {
		return nil, err
	}
	return books, nil
}
func (r *service) GetBook(id int) (Book, error) {
	book, err := r.repo.FindByID(id)

	if err != nil {
		return Book{}, err
	}
	return book, nil
}
func (r *service) CreateBook(b Book) error {
	err := r.repo.Create(b)

	if err != nil {
		return err
	}

	return nil
}
func (r *service) UpdateBook(id int, b Book) error {
	Book, err := r.repo.FindByID(id)
	if err != nil {
		return err
	}
	BookUpdate, err := Book.WithUpdatedDetails(b.Title, b.Author, b.ISBN, b.TotalCount)

	if err != nil {
		return err
	}

	err = r.repo.Update(id, BookUpdate)

	if err != nil {
		return err
	}
	return nil
}
func (r *service) DeleteBook(id int) error {
	b, err := r.repo.FindByID(id)
	if err != nil {
		return err
	}
	if b.IsBorrowed() {
		return err
	}
	err = r.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
