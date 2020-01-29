package main

import "fmt"

func main() {

	// Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d:\n", i)
	}

	//FizzBizz challenge

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBizz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Bizz")
		} else {
			fmt.Println(i)
		}

	}
}
