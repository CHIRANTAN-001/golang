package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("Present time is: ", presentTime)

	// standard time format set by golang
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.June, 28, 23, 23, 23, 0, time.UTC)

	fmt.Println("Created date is: ", createdDate)

	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 Monday"))

}