package book

import "errors"

var ErrorNotFound = errors.New("Book Not Found")
var ErrCannotDelete  = errors.New("cannot delete a borrowed book")
var ErrInvalidTotalCount = errors.New("cannot Update book")