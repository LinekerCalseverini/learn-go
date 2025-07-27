package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		b    bookstore.Book
		want int
	}
	testCases := []TestCase{
		{
			b: bookstore.Book{
				Title:  "Atlas Shrugged",
				Author: "Ayn Rand",
				Copies: 4,
			},
			want: 3,
		},
		{
			b: bookstore.Book{
				Title:  "Anatomy of the State",
				Author: "Murray A. Rothbard",
				Copies: 9,
			},
			want: 8,
		},
	}
	for _, tc := range testCases {
		result, err := bookstore.Buy(tc.b)
		if err != nil {
			t.Fatal(err)
		}
		got := result.Copies
		if tc.want != got {
			t.Errorf("bookstore.Buy(%v): want %v, got %v", tc.b, tc.want, got)
		}
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Anatomy of the State",
		Author: "Murray A. Rothbard",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	for range 100 {
		got := catalog.GetAllBooks()
		sort.Slice(got, func(i, j int) bool {
			return got[i].ID < got[j].ID
		})
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	type testCase struct {
		id   int
		want bookstore.Book
	}
	testCases := []testCase{
		{id: 1, want: bookstore.Book{ID: 1, Title: "For the Love of Go"}},
		{id: 2, want: bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}},
	}
	for _, tc := range testCases {
		got, err := catalog.GetBook(tc.id)
		if err != nil {
			t.Fatalf("TestGetBook(..., %v): %v", tc.id, err)
		}
		if !cmp.Equal(tc.want, got) {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}

func TestGetBookNotExists(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	id := 3
	_, err := catalog.GetBook(id)
	if err == nil {
		t.Errorf(
			"TestGetBookNotExists(..., %v): no error return when a book doesn't exist",
			id,
		)
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := book.NetPriceCents()
	if want != got {
		t.Errorf("book.NetPriceCents(%v): want %v, got %v", book, want, got)
	}
}
