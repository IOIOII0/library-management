package book

type Book struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	ISBN           string `json:"isbn"`
	TotalCount     int    `json:"total_count"`
	AvailableCount int    `json:"available_count"`
}

func (b Book) IsBorrowed() bool {
	return b.AvailableCount < b.TotalCount
}

func (b Book) WithUpdatedDetails(title, author, isbn string, totalCount int) (Book, error) {
        borrowed := b.TotalCount - b.AvailableCount

        if totalCount < borrowed {
                return Book{}, ErrInvalidTotalCount
        }

        return Book{
                ID:             b.ID,
                Title:          title,
                Author:         author,
                ISBN:           isbn,
                TotalCount:     totalCount,
                AvailableCount: totalCount - borrowed,
        }, nil
  }

