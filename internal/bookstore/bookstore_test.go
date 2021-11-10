package bookstore_test

import (
	"booklib/internal/bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "The Witcher - Last wish",
		Author: "Andrzej Sapkowski",
		Copies: 3,
	}
}
