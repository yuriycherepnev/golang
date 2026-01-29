package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	aloha := createAnyUser("user")

	if num, ok := aloha.(Person); ok {
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
}

func createAnyUser(role string) interface{} {
	if role == "admin" {
		return 22
	}
	return struct {
		Name string
		Age  int
	}{
		Name: "Helen",
		Age:  25,
	}
}
