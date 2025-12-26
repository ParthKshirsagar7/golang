# Notes on Pointers in Go

## Key Points
- A **pointer** is a variable that stores the memory address of another variable.
- Pointers in Go are denoted by the `*` operator (to declare or dereference) and the `&` operator (to get the address of a variable).
- Go does not support pointer arithmetic, making it safer than languages like C.
- Pointers are used to:
  - Avoid copying large data structures.
  - Allow functions to modify variables passed as arguments.

## Examples

### Declaring and Using Pointers
```go
num := 10
ptr := &num // Get the address of num
fmt.Println(ptr)  // Output: memory address of num
fmt.Println(*ptr) // Output: 10 (dereferencing the pointer)
```

### Passing by Value vs Passing by Reference
- **Pass by Value**: The function gets a copy of the variable, so changes do not affect the original variable.
```go
func changeNum(num int) {
    num = 5678
}

num := 10
changeNum(num)
fmt.Println(num) // Output: 10 (unchanged)
```

- **Pass by Reference**: The function gets a pointer to the variable, so changes affect the original variable.
```go
func changeNumInMemory(num *int) {
    *num = 1234
}

num := 10
changeNumInMemory(&num)
fmt.Println(num) // Output: 1234 (changed)
```

### Zero Value of a Pointer
- The zero value of a pointer is `nil`.
```go
var ptr *int
fmt.Println(ptr) // Output: <nil>
```

## Additional Notes
- Pointers are commonly used in Go for:
  - Modifying function arguments.
  - Managing large data structures efficiently.
  - Working with structs and methods.
- Dereferencing a `nil` pointer causes a runtime panic.
- Go's garbage collector automatically manages memory, so you don't need to manually allocate or free memory.