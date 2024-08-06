package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	service services.LibraryManager
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		service: services.NewLibraryService(),
	}
}

func (lc *LibraryController) AddBook(scanner *bufio.Scanner) {
	var book models.Book
	fmt.Print("Enter Book ID: ")
	scanner.Scan()
	book.ID, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Enter Book Title: ")
	scanner.Scan()
	book.Title = scanner.Text()
	fmt.Print("Enter Book Author: ")
	scanner.Scan()
	book.Author = scanner.Text()
	book.Status = "Available"
	lc.service.AddBook(book)
	fmt.Println("Book added successfully.")
}

func (lc *LibraryController) RemoveBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to remove: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())
	err := lc.service.RemoveBook(bookID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func (lc *LibraryController) BorrowBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to borrow: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())
	err := lc.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func (lc *LibraryController) ReturnBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to return: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())
	err := lc.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.service.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) ListBorrowedBooks(scanner *bufio.Scanner) {
	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())
	books := lc.service.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
