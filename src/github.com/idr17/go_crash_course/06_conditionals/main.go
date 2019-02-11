package main

import "fmt"

func main() {

	a := 5
	b := 10

	if a < 10 {
		fmt.Printf("%d is less than %d\n", a, b)
	} else {
		fmt.Printf("%d is less than %d\n", b, a)
	}

	color := "green"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT red or blue")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT red or blue")
	}
}
