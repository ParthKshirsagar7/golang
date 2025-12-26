# Constants

Constants are values that are fixed at compile-time and cannot be changed during the execution of the program.

1. Basic declaration
```go
const name = "parth"
```
OR
```go
const (
    a = 1
    b = 2
    c = 3
)
```

2. Untyped vs Typed constants
- **Untyped:** When you don't specify a type, the constant has no "hard" type yet. It possesses a "default type" but can be used flexibly with different types without explicit conversion.
- **Typed:** If you specify the type, the constant is strictly bound to it.
```go
const untypedInt = 42    // Untyped (default type int)
const typedInt int = 42  // Typed

var myFloat float64 = untypedInt // Works: untyped converts automatically
// var myFloat float64 = typedInt // Fails: cannot use int as float64
```

3. Rules and constratins
- You cannot assign the result of a function call to a constant. The value must be known at compile time.
- Once declared, value of a constant cannot be changed.