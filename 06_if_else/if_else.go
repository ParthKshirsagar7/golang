package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	if score := 10; score >= 90 {
		fmt.Println("Good score")
	} else {
		fmt.Println("You need to improve")
	}

	username := "parth"

	if username == "admin" {
		fmt.Println("Welcome admin, you are authorized")
	} else if username == "parth" {
		fmt.Println("Welcome parth, you are authorized")
	} else {
		fmt.Println("Unauthorized")
	}
}
