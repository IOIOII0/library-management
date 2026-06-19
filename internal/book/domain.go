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

// func (b Book) WithUpdatedDetails(borrow int) bool {
// 	if b.AvailableCount+borrow <= b.TotalCount {
// 		b.AvailableCount += borrow
// 		return true
// 	}
// 	return false
// }
