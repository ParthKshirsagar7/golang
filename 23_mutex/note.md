# Notes on Mutex in Go

## Key Points
- A **Mutex** (short for mutual exclusion) is a synchronization primitive used to prevent race conditions when multiple goroutines access shared data.
- The `sync.Mutex` type in Go provides two main methods:
  - `Lock()`: Acquires the lock, blocking other goroutines from accessing the critical section.
  - `Unlock()`: Releases the lock, allowing other goroutines to proceed.
- Mutexes ensure that only one goroutine can access the critical section at a time.

## Examples

### Using Mutex to Protect Shared Data
```go
type post struct {
    views int
    mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
    defer func() {
        p.mu.Unlock()
        wg.Done()
    }()

    p.mu.Lock()
    p.views++
}

func main() {
    var wg sync.WaitGroup
    myPost := post{views: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg)
    }

    wg.Wait()
    fmt.Println(myPost.views) // Output: 100
}
```
- **Explanation**:
  - `p.mu.Lock()`: Acquires the lock before incrementing `views`.
  - `p.mu.Unlock()`: Releases the lock after the increment is complete.
  - `wg.Done()`: Signals that the goroutine is finished.

### Key Points to Remember
- Always pair `Lock()` with `Unlock()` to avoid deadlocks.
- Use `defer` to ensure `Unlock()` is called even if the function exits early.
- Mutexes are suitable for protecting small critical sections. For more complex synchronization, consider using channels.

## Additional Notes
- **Race Conditions**: Occur when multiple goroutines access shared data simultaneously without proper synchronization.
- **Deadlocks**: Happen when a goroutine waits indefinitely for a lock that is never released.
- **Best Practices**:
  - Minimize the scope of the critical section to reduce contention.
  - Avoid holding a lock for long periods.
  - Use `sync.RWMutex` for read-heavy workloads, which allows multiple readers but only one writer.

### Read-Write Mutex
- The `sync.RWMutex` type provides separate locks for reading and writing:
  - `RLock()`: Acquires a read lock, allowing multiple readers.
  - `RUnlock()`: Releases the read lock.
  - `Lock()` and `Unlock()`: Used for writing.
```go
type counter struct {
    value int
    mu    sync.RWMutex
}

func (c *counter) read() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.value
}

func (c *counter) write(val int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value = val
}
```
- Use `RWMutex` when reads are frequent and writes are rare.