package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
func main() {
	// uninitialized slice is nil
	var nums []int

	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var arr = make([]int, 0, 5)
	// capacity -> maximum number of elements that can fit currently
	fmt.Println(cap(arr)) // 5
	fmt.Println(arr==nil) // false
	arr = append(arr, 1, 2, 3, 4, 5, 6)
	fmt.Println(cap(arr)) // 10
	fmt.Println(arr) // [1 2 3 4 5 6]

	var arr2 = make([]int, len(arr))
	arr = append(arr, 7)
	// copy function

	copy(arr2, arr)

	fmt.Println(arr, arr2) // [1 2 3 4 5 6 7] [1 2 3 4 5 6]

	// slice operator
	var arr3 = []int{1, 2, 3}

	fmt.Println(arr3[0:1]) // [1]
	fmt.Println(arr3[:2]) // [1 2]
	fmt.Println(arr3[1:]) // [2 3]

	var num1 = []int{1, 2}
	var num2 = []int{1, 2}
	fmt.Println(slices.Equal(num1, num2)) // true

	// 2d slice
	var twoDArr = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoDArr)
}
