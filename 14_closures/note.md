# Notes on Closures in Go

## Key Points
- A **closure** is a function value that references variables from outside its body.
- In Go, closures are created when a function is defined inside another function and accesses variables from the outer function.
- Closures allow you to encapsulate state and behavior together.
- Variables captured by a closure are bound to the closure, meaning their values persist between function calls.

## Examples

### Creating a Closure
- The inner function `func() int` captures the variable `count` from the outer function `counter`.
```go
func counter() func() int {
    var count int = 0

    return func() int {
        count += 1
        return count
    }
}

increment := counter()
fmt.Println(increment()) // Output: 1
fmt.Println(increment()) // Output: 2
```

### Explanation
- Each call to `counter` creates a new instance of `count`.
- The `increment` function retains access to the `count` variable even after `counter` has finished executing.

### Using Closures for State Management
- Closures are useful for managing state in a clean and concise way.
```go
func adder(base int) func(int) int {
    return func(value int) int {
        base += value
        return base
    }
}

add := adder(10)
fmt.Println(add(5))  // Output: 15
fmt.Println(add(10)) // Output: 25
```

## Additional Notes
- Closures are commonly used in Go for:
  - Creating function factories.
  - Managing state without using global variables.
  - Defining inline functions with access to outer variables.
- Captured variables are stored on the heap, so they persist as long as the closure exists.
- Be cautious when using closures in concurrent code, as shared variables may lead to race conditions.