# Notes on Generics in Go

## Key Points
- **Generics** allow you to write reusable and type-safe code by defining functions and types that work with any data type.
- Introduced in Go 1.18, generics use type parameters to enable flexibility while maintaining type safety.
- Type parameters are specified using square brackets `[]` and can have constraints.
- Constraints define the permissible types for a type parameter (e.g., `comparable`, `any`).

## Examples

### Generic Functions
- A generic function works with any type specified as a type parameter.
```go
func printSlice[T comparable](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

vals := []bool{true, false, true}
printSlice(vals) // Output: true false true
```
- `T` is the type parameter, and `comparable` is the constraint, ensuring that `T` supports comparison operations.

### Generic Types
- Generic types allow you to define data structures that work with any type.
```go
type stack[T any] struct {
    elements []T
}

myStack := stack[int]{
    elements: []int{1, 2, 3},
}
fmt.Println(myStack) // Output: { [1 2 3] }
```
- `T` is the type parameter, and `any` is the constraint, meaning `T` can be any type.

### Constraints
- Constraints restrict the types that can be used as type parameters.
- Common constraints:
  - `any`: Allows any type.
  - `comparable`: Allows types that support comparison operations (e.g., `==`, `!=`).
- Custom constraints can be defined using interfaces.
```go
type Number interface {
    int | float64
}

func sum[T Number](a, b T) T {
    return a + b
}

fmt.Println(sum(10, 20))       // Output: 30
fmt.Println(sum(3.14, 2.71))  // Output: 5.85
```

## Additional Notes
- Generics improve code reusability and reduce duplication.
- Use constraints to ensure type safety and define valid operations for type parameters.
- Generics are particularly useful for collections, algorithms, and data structures.
- Avoid overusing generics when simpler solutions (e.g., concrete types) suffice, as they can make code harder to read.