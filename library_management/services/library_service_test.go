package services

import (
	"library_management/models"
	"testing"
)

func TestLibraryService(t *testing.T) {
	library := NewLibraryService()

	book1 := models.Book{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Status: "Available"}
	book2 := models.Book{ID: 2, Title: "Learning Go", Author: "Jon Bodner", Status: "Available"}
	library.AddBook(book1)
	library.AddBook(book2)

	if len(library.ListAvailableBooks()) != 2 {
		t.Errorf("expected 2 available books, got %d", len(library.ListAvailableBooks()))
	}

	err := library.RemoveBook(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(library.ListAvailableBooks()) != 1 {
		t.Errorf("expected 1 available book, got %d", len(library.ListAvailableBooks()))
	}

	
	memberID := 1001
	err = library.BorrowBook(1, memberID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(library.ListAvailableBooks()) != 0 {
		t.Errorf("expected 0 available books, got %d", len(library.ListAvailableBooks()))
	}

	borrowedBooks := library.ListBorrowedBooks(memberID)
	if len(borrowedBooks) != 1 {
		t.Errorf("expected 1 borrowed book, got %d", len(borrowedBooks))
	}

	// Test ReturnBook
	err = library.ReturnBook(1, memberID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(library.ListAvailableBooks()) != 1 {
		t.Errorf("expected 1 available book, got %d", len(library.ListAvailableBooks()))
	}

	borrowedBooks = library.ListBorrowedBooks(memberID)
	if len(borrowedBooks) != 0 {
		t.Errorf("expected 0 borrowed books, got %d", len(borrowedBooks))
	}

	// Test BorrowBook with non-existent book
	err = library.BorrowBook(3, memberID)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	// Test ReturnBook with non-existent book
	err = library.ReturnBook(3, memberID)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	// Test RemoveBook with non-existent book
	err = library.RemoveBook(3)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
