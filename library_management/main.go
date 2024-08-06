package main

import (
	"bufio"
	"fmt"
	"os"
	"library_management/controllers"
)

func main() {
	controller := controllers.NewLibraryController()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			controller.AddBook(scanner)
		case "2":
			controller.RemoveBook(scanner)
		case "3":
			controller.BorrowBook(scanner)
		case "4":
			controller.ReturnBook(scanner)
		case "5":
			controller.ListAvailableBooks()
		case "6":
			controller.ListBorrowedBooks(scanner)
		case "7":
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
