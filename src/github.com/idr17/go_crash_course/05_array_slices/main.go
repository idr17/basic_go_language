package main

import "fmt"

func main() {

	// 1 declare and assign
	// var fruitArr = [2]string{"Orange", "Apple"}
	// fruitArr[0]

	// 2 declare and assign
	fruitArr := []string{"Orange", "Apple", "Grape", "Pineaple"}

	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:3])
}
