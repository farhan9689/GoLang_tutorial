package main

import "fmt"

func main() {
	// Arraye

	// var fruitArr [2]string
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	//Declare and assign  (another way of declaring and assigning)

	fruitArr := [2]string{"apple", "orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//slice Example

	fruitSlice := []string{"apple", "banana", "guawa", "cherry"} //if we use slice we dont have t

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) // this is a "range" method to print something in range(it will print from index 1 to 2 for range 1:3 i.e index 3 is excluded)
}
