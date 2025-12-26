package main

import "fmt"

// pass by value
func changeNum(num int) {
	num = 5678
}

// pass by reference
func changeNumInMemory(num *int) {
	*num = 1234
}

func main() {
	num := 10

	fmt.Println("Before calling changeNum:", num)
	changeNum(num)
	fmt.Println("After calling changeNum:", num)

	fmt.Println("Before calling changeNumInMemory:", num)
	changeNumInMemory(&num)
	fmt.Println("After calling changeNumInMemory:", num)
}
