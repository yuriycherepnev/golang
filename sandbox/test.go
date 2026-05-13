package main

import "fmt"

type User struct {
	count int
	book  Book
}

type Book struct {
	count int
}

type UserInterface interface {
	Inc()
}

func main() {
	user := User{
		count: 2,
		book:  Book{count: 2},
	}

	user.book.Inc()

	fmt.Println(user.book.count) // 1
}

func (u *User) Inc() {
	u.count++
}

func (u *Book) Inc() {
	u.count++
}
