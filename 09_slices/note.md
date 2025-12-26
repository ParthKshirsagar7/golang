# Notes on Slices in Go

## Key Points
- **Slices** are a dynamic, flexible view into the elements of an array.
- Unlike arrays, slices are not fixed in size and can grow or shrink.
- A slice is a descriptor of an array segment and consists of three components:
  - **Pointer**: Points to the underlying array.
  - **Length**: The number of elements in the slice.
  - **Capacity**: The maximum number of elements the slice can grow to without reallocating.
- Slices are the most commonly used construct in Go for working with collections.

## Examples

### Declaring and Initializing Slices
```go
var nums []int
fmt.Println(nums == nil) // true
fmt.Println(len(nums))   // 0

var arr = make([]int, 0, 5)
fmt.Println(cap(arr)) // 5
fmt.Println(arr == nil) // false
```

### Appending to Slices
```go
arr = append(arr, 1, 2, 3, 4, 5, 6)
fmt.Println(cap(arr)) // 10 (capacity doubles when exceeded)
fmt.Println(arr)      // [1 2 3 4 5 6]
```

### Copying Slices
```go
var arr2 = make([]int, len(arr))
copy(arr2, arr)
fmt.Println(arr, arr2) // [1 2 3 4 5 6 7] [1 2 3 4 5 6]
```

### Slice Operator
```go
var arr3 = []int{1, 2, 3}
fmt.Println(arr3[0:1]) // [1]
fmt.Println(arr3[:2])  // [1 2]
fmt.Println(arr3[1:])  // [2 3]
```

### Comparing Slices
- Use the `slices.Equal` function from the `slices` package to compare slices.
```go
var num1 = []int{1, 2}
var num2 = []int{1, 2}
fmt.Println(slices.Equal(num1, num2)) // true
```

### Multi-dimensional Slices
```go
var twoDArr = [][]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println(twoDArr) // [[1 2 3] [4 5 6]]
```

## Additional Notes
- Slices are backed by arrays, and modifying a slice modifies the underlying array.
- The `append` function may allocate a new array if the capacity is exceeded, leaving the original array unchanged.
- Use the `make` function to create slices with a specified length and capacity.
- Slices are nil by default if uninitialized.