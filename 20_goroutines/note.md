# Notes on Goroutines in Go

## Key Points
- **Goroutines** are lightweight threads managed by the Go runtime.
- They enable concurrent execution of functions, making Go highly efficient for concurrent programming.
- Use the `go` keyword to start a new goroutine.
- Goroutines are non-blocking and run independently of the main program flow.
- The Go runtime schedules goroutines across available CPU cores.

## Examples

### Starting a Goroutine
- Use the `go` keyword to execute a function as a goroutine.
```go
func task(id int) {
    fmt.Println("Doing task", id)
}

go task(1) // Starts the task function as a goroutine
```

### Anonymous Functions as Goroutines
- You can use anonymous functions to pass arguments to goroutines.
```go
for i := 0; i < 10; i++ {
    go func(i int) {
        fmt.Println(i)
    }(i)
}
```
- The argument `i` is passed explicitly to avoid capturing the loop variable.

## Additional Notes
- Goroutines are extremely lightweight compared to OS threads, with a small initial memory footprint (~2 KB).
- The Go runtime dynamically grows and shrinks the stack size of goroutines as needed.
- Be cautious of race conditions when accessing shared data. Use synchronization primitives (e.g., `sync.Mutex`) or channels to avoid data races, which will be discussed in separate notes.
- Goroutines are a fundamental building block for concurrent programming in Go, enabling scalable and efficient applications.