package main

import "fmt"

func greeting() func(string) string { // what is this technique? currently dont know

	return func(s string) string {
		return "hello" + s
	}
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	show := greeting() // didn't understand
	fmt.Println(show("Farhan"))
	fmt.Println(getSum(2, 5))
}
