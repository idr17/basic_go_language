package main

import "fmt"

func main() {

	// 1
	// declare
	// emails := make(map[string]string)

	// // assign key value
	// emails["indro"] = "indera.allezz@gmail.com"
	// emails["eko"] = "eko@gmail.com"
	// emails["jon"] = "jon@gmail.com"

	// 2 declare and assign kv
	emails := map[string]string{"indro": "indera.allezz@gmail.com", "eko": "eko@gmail.com", "jon": "jon@gmail.com"}

	// delete
	delete(emails, "jon")

	// add
	emails["new"] = "new@gmail.com"

	fmt.Println(emails)
}
