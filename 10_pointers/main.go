package main

import "fmt"

func main() {
	a := 5
	b := &a // memory addres

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // * represents pointer

	// use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
