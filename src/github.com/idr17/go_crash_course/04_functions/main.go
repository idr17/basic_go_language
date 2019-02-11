package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sum(var1, var2 int) int {
	return var1 + var2
}

func main() {
	fmt.Println(greeting("Boys.. "))
	fmt.Println(sum(3, 6))
}
