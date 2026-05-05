package main

import (
	"fmt"
	"hash/maphash"
	"unsafe"
)

type seedStruct struct {
	s uint64
}

func main() {
	var seed maphash.Seed
	seed = maphash.MakeSeed()

	s := (*seedStruct)(unsafe.Pointer(&seed))
	fmt.Println(s.s)

	var seed1 maphash.Seed
	seed1 = maphash.MakeSeed()

	s1 := (*seedStruct)(unsafe.Pointer(&seed1))
	fmt.Println(s1.s)
}
