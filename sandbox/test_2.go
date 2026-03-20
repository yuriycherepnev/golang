package main

import "fmt"

func main() {
	a := "hello"
	b := a[0:2]
	fmt.Println(&a, a)
	fmt.Println(&b, b)

	a = "world"
	fmt.Println(&a, a)
	fmt.Println(&b, b)

	s := "hАllo"

	for _, v := range s {
		fmt.Println(string(v))
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

}
