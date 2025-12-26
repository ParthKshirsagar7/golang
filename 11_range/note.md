# Notes on `range` in Go

## Key Points
- The `range` keyword in Go is used to iterate over elements in various data structures such as slices, arrays, maps, strings, and channels.
- It provides a concise way to loop through collections.
- Depending on the data structure, `range` returns different values:
  - **Slices/Arrays**: Index and value.
  - **Maps**: Key and value.
  - **Strings**: Index and Unicode code point (rune).
  - **Channels**: Value sent on the channel.

## Examples

### Iterating Over a Slice
```go
nums := []int{1, 2, 3, 4, 5}
for i, v := range nums {
    fmt.Printf("nums[%d] = %d\n", i, v)
}
// Output:
// nums[0] = 1
// nums[1] = 2
// nums[2] = 3
// nums[3] = 4
// nums[4] = 5
```

### Iterating Over a Map
```go
m := map[string]string{"first_name": "Parth", "last_name": "Kshirsagar"}
for key, value := range m {
    fmt.Println(key, value)
}
// Output:
// first_name Parth
// last_name Kshirsagar
```

### Iterating Over a String
- The `range` keyword iterates over Unicode code points in a string.
- The index represents the starting byte of the rune.
```go
for i, c := range "😂golang" {
    fmt.Println(i, c)
}
// Output:
// 0 128514
// 4 103
// 5 111
// 6 108
// 7 97
// 8 110
// 9 103
```

### Ignoring Values
- Use the blank identifier `_` to ignore values you don't need.
```go
for _, v := range nums {
    fmt.Println(v) // Only prints values
}

for key := range m {
    fmt.Println(key) // Only prints keys
}
```

## Additional Notes
- When iterating over a map, the order of iteration is not guaranteed and can vary.
- The `range` keyword is versatile and simplifies iteration over collections.
- For strings, the Unicode code point can be converted to a character using `string(c)`.
- Use `range` with channels to receive values until the channel is closed.
