package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Prog"
	Book1.author = "Tugay"
	Book1.subject = "tarih"
	Book1.book_id = 1

	fmt.Printf("Book title : %s\n", Book1.title)
	fmt.Printf("Book author : %s\n", Book1.author)
	fmt.Printf("Book subject : %s\n", Book1.subject)
	fmt.Printf("Book book_id : %d\n", Book1.book_id)

	Book2.title = "python"
	Book2.author = "tugay2"
	Book2.subject = "deneme"
	Book2.book_id = 2

	pringBook(Book2)
}

func pringBook(books Books)  {
	fmt.Printf("Book title : %s\n", books.title)
	fmt.Printf("Book author : %s\n", books.author)
	fmt.Printf("Book subject : %s\n", books.subject)
	fmt.Printf("Book book_id : %d\n", books.book_id)
}
