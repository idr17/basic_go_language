package main

import "fmt"

func main() {

	a := 5
	b := &a // means memory location of pointer a

	fmt.Println(a)
	fmt.Printf("Type of a %T\n", a)
	fmt.Printf("Type of b %T\n", b) // return *int => means pointer location memory and value stored is int

	*b = 10 // change value from pointer address
	fmt.Println(a)
}
