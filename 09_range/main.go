package main

import "fmt"

func main() {

	ids := []int{33, 75, 15, 3, 42, 9}

	for i, id := range ids {
		fmt.Printf("%d - ID:%d\n", i, id)

	}

	//Without using index

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)

	}

	// for id := range ids {             another way of writing above code
	// 	fmt.Printf("ID: %d\n", id)

	// }
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	// range with maps

	emails := map[string]string{
		"farhan": "farhan@gmail.com",
		"mirza":  "mirza@gmail.com",
		"bob":    "bob@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}
	// printing just the values
	for _, v := range emails { //if use any variable inplace of "_," than we get error
		fmt.Printf("email : %s\n", v)
	}
	// pring just the keys

	for k := range emails {
		fmt.Printf(" Name: %s \n", k)
	}
}
