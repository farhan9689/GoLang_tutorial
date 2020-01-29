package main

import "fmt"

func main() {

	a := 5

	b := &a

	fmt.Println(a, b)

	// use * to red value

	fmt.Println(*b)  //o/p -> 5
	fmt.Println(*&a) //o/p -> 5

	//change value with pointer

	*b = 10
	fmt.Println(a) //printing updated value
}
