package main

import "fmt"

func main() {

	// 1
	// declare. map is key value pairs
	// emails := make(map[string]string) // map[string] => key string; last string => value string

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

	fmt.Printf("%T", emails)
}
