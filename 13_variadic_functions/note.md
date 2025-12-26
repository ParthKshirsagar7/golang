# Notes on Variadic Functions in Go

## Key Points
- **Variadic functions** in Go allow you to pass a variable number of arguments to a function.
- The syntax for a variadic parameter is `...type`, where `type` is the type of the arguments.
- A variadic parameter must be the last parameter in the function signature.
- Variadic arguments are received as a slice within the function.

## Examples

### Defining and Using Variadic Functions
- A variadic function can accept zero or more arguments of the specified type.
```go
func sum(nums ...int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}

nums := []int{1, 2, 3, 4, 5}
result := sum(nums...)
fmt.Println(result) // Output: 15
```

### Passing Arguments
- You can pass individual arguments or a slice using the `...` operator.
```go
fmt.Println(sum(1, 2, 3)) // Output: 6

nums := []int{4, 5, 6}
fmt.Println(sum(nums...)) // Output: 15
```

### Variadic Functions with Other Parameters
- A variadic parameter must always be the last parameter in the function signature.
```go
func printMessage(message string, nums ...int) {
    fmt.Println(message)
    for _, v := range nums {
        fmt.Println(v)
    }
}

printMessage("Numbers:", 1, 2, 3)
// Output:
// Numbers:
// 1
// 2
// 3
```

## Additional Notes
- Variadic functions are useful for operations like summing numbers, formatting strings, or handling a dynamic number of inputs.
- If no arguments are passed to a variadic function, the variadic parameter is treated as an empty slice.
- Use the `...` operator to expand a slice into individual arguments when calling a variadic function.