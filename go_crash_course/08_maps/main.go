package main

import "fmt"

func main() {
	// MAPS
	// Define Maps

	// emails := make(map[string]string)
	emails := map[string]string{
		"Banna": "banna@gmail.com",
		"Kapil": "kapil@gmail.com"}

	// Assign Key Values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
