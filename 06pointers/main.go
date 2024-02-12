package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int 
	// fmt.Println("The value of ptr is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("The value of ptr is: ", ptr)
	fmt.Println("The value of ptr is: ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("new value is: ", myNumber)
}