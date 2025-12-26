package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	} 

	// type switch
	var k interface{} = false
	switch i := k.(type) {
	case int:
		fmt.Println("Its an integer")
	case string:
		fmt.Println("Its a string")
	case bool:
		fmt.Println("Its a boolean")
	default:
		fmt.Println("Other", i)
	}

	// fallthrough
	switch num := 2; num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three") // This will print because of fallthrough
	}
	// Output: Two, Three
}
