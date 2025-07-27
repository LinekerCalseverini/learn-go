package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, book := range c {
		result = append(result, book)
	}
	return result
}

func (c Catalog) GetBook(id int) (Book, error) {
	book, ok := c[id]
	if !ok {
		return book, fmt.Errorf("no book with id %d found", id)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	return b.PriceCents * (100 - b.DiscountPercent) / 100
}
