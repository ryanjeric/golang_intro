package main

import "fmt"

func main() {
	// Main types

	// string
	// bool
	// in
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	// var name string = "Ryan"
	var age int32 = 10
	var isCool = false
	var size float32 = 2.3

	// shorthand
	name := "Sasa"
	//size := 1.3 // float64

	//name,email := "sasa","testemail" # single line

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T", size) //
}
