package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// Assign keyvalues
	/* emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com" */

	// Declare map and kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)

	fmt.Println(emails["Bob"])
	fmt.Println(emails["Sharon"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
