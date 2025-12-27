# Notes on Wait Groups in Go

## Key Points
- **Wait Groups** are part of the `sync` package and are used to wait for a collection of goroutines to complete.
- A `sync.WaitGroup` provides a simple way to synchronize multiple goroutines.
- The `WaitGroup` has three main methods:
  - `Add(delta int)`: Increments the counter by `delta`.
  - `Done()`: Decrements the counter by 1 (usually called with `defer`).
  - `Wait()`: Blocks until the counter becomes zero.

## Examples

### Using Wait Groups
```go
func task(id int, w *sync.WaitGroup) {
    defer w.Done() // Decrement the counter when the goroutine completes
    fmt.Println("Doing task:", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i <= 10; i++ {
        wg.Add(1) // Increment the counter for each goroutine
        go task(i, &wg)
    }

    wg.Wait() // Block until all goroutines finish
}
```
- **Explanation**:
  - `wg.Add(1)`: Adds one to the counter for each new goroutine.
  - `defer wg.Done()`: Ensures the counter is decremented when the goroutine finishes.
  - `wg.Wait()`: Blocks the main goroutine until the counter reaches zero.

### Key Points to Remember
- Always pass the `WaitGroup` by reference (e.g., `*sync.WaitGroup`).
- Use `defer wg.Done()` to ensure the counter is decremented even if the goroutine panics.
- Avoid calling `wg.Add()` after calling `wg.Wait()` to prevent race conditions.

## Additional Notes
- Wait Groups are a simple and effective way to synchronize goroutines.
- They are commonly used when you need to wait for multiple goroutines to complete before proceeding.
- For more advanced synchronization, consider using channels or other primitives from the `sync` package.