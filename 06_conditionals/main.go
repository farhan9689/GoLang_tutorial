package main

import "fmt"

func main() {
	x := 15
	y := 10

	//if statement
	if x < y {
		fmt.Printf("%d is less than %d \n", x, y) // %d is kind of place holder to print tha value
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	//else if (nested)

	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red ")
	}

	// switch example

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}
}
