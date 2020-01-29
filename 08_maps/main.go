package main

import "fmt"

func main() {

	// Defining maps
	// emails := make(map[string]string)

	// // assigning key values to the map

	// emails["farhan"] = "farhan@gmail.com"
	// emails["mirza"] = "mirza@gmail.com"
	// emails["bob"] = "bob@gmail.com"

	//Defining and assigning maps at the same time

	emails := map[string]string{
		"farhan": "farhan@gmail.com",
		"mirza":  "mirza@gmail.com",
		"bob":    "bob@gmail.com"}

	emails["mike"] = "mike@gmail.com"
	fmt.Println(emails["farhan"])
	fmt.Println(len(emails)) // gives you length of map
	fmt.Println(emails)
	// Delete from map

	delete(emails, "bob")

	fmt.Println(emails)
}
