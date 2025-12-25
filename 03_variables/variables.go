package main

import "fmt"

func main() {
	// standard way
	var name string = "parth"
	var a, b, c int = 1, 2, 3

	// golang infers type of the variable
	var country = "India"

	// shorthand (walrus operator)
	age := 18

	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a)

	fmt.Println(name)
	fmt.Println(country)
	fmt.Println(age)
}
