# If-Else

In Go, if statements are straightforward.
1. Basic syntax
```go
age := 18

if age >= 18 {
    fmt.Println("Person is an adult")
} else {
    fmt.Println("Person is not an adult")
}
```

2. If with an initialization statement:
```go
if score := 10; score >= 90 {
    fmt.Println("Good score")
} else {
    fmt.Println("You need to improve")
}
```
Note that score declared here is only in score until the end of the if/else block.

3. Multiple conditions if-else:
```go
username := "parth"

if username == "admin" {
    fmt.Println("Welcome admin, you are authorized")
} else if username == "parth" {
    fmt.Println("Welcome parth, you are authorized")
} else {
    fmt.Println("Unauthorized")
}
```
