package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is greater than %d", x, y)
	}

	fmt.Println()

	// else if
	color := "red"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	/* Switch */
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}

}
