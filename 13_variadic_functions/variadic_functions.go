package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
