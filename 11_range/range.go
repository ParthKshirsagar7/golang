package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i, v := range nums {
		fmt.Printf("nums[%d] = %d\n", i , v)
	}

	m := map[string]string{"first_name": "Parth", "last_name": "Kshirsagar"}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// unicode point rune
	// i -> starting byte of rune
	for i, c := range "😂golang" {
		fmt.Println(i, c)
	}
}
