//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	title        string
	author       string
	checkedOutAt time.Time
	returnedAt   time.Time
	member       *Member
}

func (b *Book) Checkout(m *Member) {
	b.checkedOutAt = time.Now()
	b.member = m
}

func (b *Book) Return() {
	b.returnedAt = time.Now()
	b.checkedOutAt = time.Time{}
	b.member = &Member{}
}

func (b *Book) String() string {
	layout := "2006/01/02 15:04:05"
	return fmt.Sprintf("Title: %s, Author: %s, Checkout out: %s, Returned: %s", b.title, b.author, b.checkedOutAt.Format(layout), b.returnedAt.Format(layout))
}

type Member struct {
	name  string
	books []*Book
}

func (m *Member) CheckoutBook(b *Book) {
	b.Checkout(m)
	m.books = append(m.books, b)
}

func (m *Member) ReturnBook(b *Book) {
	b.Return()
	for i := range m.books {
		if m.books[i].title == b.title {
			m.books = append(m.books[:i], m.books[i+1:]...)
		}
	}
}

func (m *Member) String() string {
	s := fmt.Sprintf("Name: %s", m.name)
	for i := range m.books {
		s = fmt.Sprintf("%s\nMember Book: %s", s, m.books[i].String())
	}
	return s
}

type Library struct {
	books   []Book
	members []Member
}

func (l *Library) String() string {
	res := ""
	for i := range l.members {
		res = fmt.Sprintf("%sMember: %s \n", res, l.members[i].String())
	}
	for i := range l.books {
		res = fmt.Sprintf("%sBook: %s \n", res, l.books[i].String())
	}

	return res
}

func main() {
	members := []Member{{name: "John"}, {name: "Tom"}, {name: "Joanna"}}
	books := []Book{{title: "Pan Taduesz", author: "Henryk Sienkiewicz"}, {title: "Harry Potter: Prisoner of Azkaban", author: "J.K. Rowling"}, {title: "The Lord of the Rings", author: "Tolkien"}}
	library := &Library{books, members}
	fmt.Print(library)

	library.members[0].CheckoutBook(&library.books[0])
	fmt.Print(library)

	library.members[0].ReturnBook(&library.books[0])
	fmt.Print(library)
}
