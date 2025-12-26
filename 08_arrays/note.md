# Notes on Arrays in Go

## Key Points
- Arrays in Go have a **fixed size**.
- They are **memory optimized** and provide **constant time access** to elements.
- Default (zeroed) values for array elements depend on the type:
  - `int` -> `0`
  - `string` -> `""`
  - `bool` -> `false`

## Examples

### Declaring and Initializing Arrays
```go
var nums [4]int
nums[0] = 1
fmt.Println(nums) // Output: [1 0 0 0]

var vals [4]bool
vals[2] = true
fmt.Println(vals) // Output: [false false true false]

var names [3]string
names[0] = "parth"
fmt.Println(names) // Output: ["parth" "" ""]
```

### Using Short Declaration
```go
scores := [5]int{10, 20, 30, 40, 50}
fmt.Println(scores) // Output: [10 20 30 40 50]
```

### Multi-dimensional Arrays
```go
arr := [2][2]int{{1, 2}, {3, 4}}
fmt.Println(arr) // Output: [[1 2] [3 4]]
```