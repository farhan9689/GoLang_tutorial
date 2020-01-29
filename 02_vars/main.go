package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Creating variable using var

	// Short hand
	name := "farhan" // we can not use short hand outside the main function. But we can assign the global variable using var

	name1, email := "farhan", "farhan@gmail.com"
	size := 1.2 //default float will be float64. If you want float32 you have to assign it explicitly.
	var age = 24
	const isCool = true
	// isCool = false                  cannot reassign the const value
	fmt.Println(name, age, isCool, name1, email)
	fmt.Printf("%T\n", size)
}
