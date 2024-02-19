package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "Typescript"

	fmt.Println(languages)

	fmt.Println("JS Abbreviation is", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// loops are interesting IN GOLQANG

	for _, value := range languages {
		fmt.Printf("Value is %v", value)
	}

}