# Notes on Functions in Go

## Key Points
- Functions in Go are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.
- Go supports:
  - **Normal functions** with parameters and return values.
  - **Multiple return values**.
  - **Higher-order functions** (functions that take other functions as arguments or return functions).
  - Anonymous functions (functions without a name).

## Examples

### Normal Functions
- Functions can take parameters and return a single value.
```go
func add(a int, b int) int {
    return a + b
}

fmt.Println(add(10, 20)) // Output: 30
```

### Multiple Return Values
- Functions can return multiple values.
```go
func getLanguages() (string, string, string) {
    return "golang", "javascript", "python"
}

a, b, c := getLanguages()
fmt.Println(a, b, c) // Output: golang javascript python
```

### Higher-Order Functions
- Functions can take other functions as arguments.
```go
func higherOrder(fn func(a int) int) {
    fmt.Println(fn(1))
}

fn := func(a int) int {
    return a
}
higherOrder(fn) // Output: 1
```

### Returning Functions
- Functions can return other functions.
```go
func returnFunc() func(a int) int {
    return func(a int) int {
        return 100
    }
}

returnedFn := returnFunc()
fmt.Println(returnedFn(10)) // Output: 100
```

### Anonymous Functions
- Functions without a name can be defined inline.
```go
fn := func(a int) int {
    return a * 2
}
fmt.Println(fn(5)) // Output: 10
```

## Additional Notes
- Functions in Go can have named return values, which act as variables initialized to their zero values.
- Variadic functions allow passing a variable number of arguments (e.g., `func sum(nums ...int) int`).
- Go does not support default arguments or function overloading.
- Use `defer` to ensure a function call is executed at the end of the surrounding function's execution.