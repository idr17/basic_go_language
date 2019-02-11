package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	// fmt.Println(ids)

	// for k, id := range ids {
	// 	fmt.Println(k, id)
	// }

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("total", sum)

	emails := map[string]string{"indro": "indera.allezz@gmail.com", "eko": "eko@gmail.com", "jon": "jon@gmail.com"}
	for k := range emails {
		fmt.Println("name:", k)
	}
}
