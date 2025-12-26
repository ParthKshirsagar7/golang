# Notes on Maps in Go

## Key Points
- **Maps** in Go are collections of key-value pairs, where keys are unique.
- Maps are reference types and must be initialized before use.
- The zero value of a map is `nil`, and a `nil` map behaves like an empty map when reading but causes a runtime panic when writing.
- Maps provide **constant time complexity** for basic operations like adding, deleting, and retrieving elements.

## Examples

### Creating and Initializing Maps
```go
// Using make function
m := make(map[string]string)

// Initializing while declaration
mp := map[string]int{"price": 40, "phones": 3}
fmt.Println(mp) // Output: map[phones:3 price:40]
```

### Adding and Accessing Elements
```go
m["name"] = "golang"
m["area"] = "backend"
fmt.Println(m["name"], m["area"]) // Output: golang backend
```

### Default Values for Non-existent Keys
- Accessing a key that does not exist in the map returns the zero value of the value type.
```go
info := make(map[string]int)
info["age"] = 18
fmt.Println(info["age"])       // Output: 18
fmt.Println(info["something"]) // Output: 0
```

### Checking Key Existence
- Use the **comma ok idiom** to check if a key exists in the map.
```go
val, ok := mp["price"]
fmt.Println(val) // Output: 40
fmt.Println(ok)  // Output: true
```

### Deleting Elements
```go
delete(info, "score")
fmt.Println(info) // Output: map[age:18]
```

### Clearing a Map
- Use a custom `clear` function to remove all elements from a map.
```go
clear(info)
fmt.Println(info) // Output: map[]
```

### Comparing Maps
- Use the `maps.Equal` function from the `maps` package to compare two maps.
```go
m1 := map[string]int{"price": 40, "phones": 3}
m2 := map[string]int{"price": 40, "phones": 3}
fmt.Println(maps.Equal(m1, m2)) // Output: true
```

## Additional Notes
- Maps are not safe for concurrent use. Use synchronization mechanisms like `sync.Mutex` or `sync.Map` for concurrent access.
- The `len` function returns the number of key-value pairs in the map.
- Maps are highly versatile and are commonly used for lookups, caching, and grouping data.