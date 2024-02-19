package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// no inheritance in golang: No super or parent

	john := User{"john", "john@email.com", true, 20 }

	fmt.Println(john)

	// +%v -> get all data from the struct with field names
	fmt.Printf("John's details are %+v\n", john)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


