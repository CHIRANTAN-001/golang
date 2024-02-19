package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Apple", "Orange", "Banana", "mango", "papaya")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 254
	highScores[1] = 354
	highScores[2] = 454
	highScores[3] = 554

	highScores = append(highScores, 654, 754, 954, 854)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// ho9w to remove a value from slices based on index

	var courses = []string{"reactjs", "golang", "js", "express", "rust"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}	
