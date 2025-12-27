package main

import "fmt"

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// LIFO
type stack[T any] struct {
	elements []T
}

// func (s stack)

func main() {
	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript", "javascript", "python", "java"}
	vals := []bool{true, false, true}
	printSlice(vals)

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
}
