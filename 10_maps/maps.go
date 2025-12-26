package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// getting an element
	fmt.Println(m["name"], m["area"])

	// if key value doesn't exist in the map, it returns zero value
	info := make(map[string]int)
	info["age"] = 18
	fmt.Println(info["age"]) // 18
	fmt.Println(info["something"]) // 0

	info["score"] = 100
	fmt.Println(len(info)) // 2
	fmt.Println(info) // map[age:18 score:100]
	delete(info, "score")
	fmt.Println(info) // map[age:18]

	clear(info)
	fmt.Println(info) // map[]

	// initializing while declaration
	mp := map[string]int{"price": 40, "phones": 3}
	fmt.Println(mp) // map[phones:3 price:40]

	val, ok := mp["price"]
	fmt.Println(val) // 40
	fmt.Println(ok) // true

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}

	fmt.Println(maps.Equal(m1, m2)) // true
}
