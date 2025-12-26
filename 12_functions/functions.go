package main

import "fmt"

func main() {
	// normal function
	fmt.Println(add(10,20))

	// multiple return values
	a, b, c := getLanguages()
	fmt.Println(a, b, c)

	// higher order function
	fn := func (a int) int {
		return a
	}
	higherOrder(fn)

	// return value function
	returnedFn := returnFunc()
	fmt.Println(returnedFn(10))
}

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "python"
}

func higherOrder(fn func(a int) int){
	fmt.Println(fn(1))
}

func returnFunc() func(a int) int {
	return func (a int) int {
		return 100
	}
}