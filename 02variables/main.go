package main


import "fmt"

const JwtSecretKey string = "jwt-secret-key" //public

func main() {
	var username string = "Chirantan"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n", username)

	fmt.Println("==========================")

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)

	fmt.Println("==========================")

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)

	fmt.Println("==========================")

	var smallFloat float64 = 255.45554548684686846846464
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n", smallFloat)

	fmt.Println("==========================")

	//default values and some aliases

	var anotherVariable float64
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type: %T \n", anotherVariable)

	fmt.Println("==========================")

	//implicit type

	var website = "http://chirantan.com"
	fmt.Println("Website: ", website)

	//no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println("==========================")

	fmt.Println(JwtSecretKey)
	fmt.Printf("Variables is of type: %T \n", JwtSecretKey)

}
