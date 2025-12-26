package main

import "fmt"

func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false

	var nums [4]int
	nums[0] = 1
	fmt.Println(nums)

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	var names [3]string
	names[0] = "parth"
	fmt.Println(names)

	// var scores [5]int = [5]int{10, 20, 30, 40, 50}
	scores := [5]int{10, 20, 30, 40, 50}
	fmt.Println(scores)

	// 2d arrays
	arr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr)

	// about arrays
	// - fixed size
	// - memory optimization
	// - constant time access
}