# Notes on Channels in Go

## Key Points
- **Channels** are a way to communicate between goroutines in Go.
- They allow goroutines to send and receive data, enabling synchronization and coordination.
- Channels are created using the `make` function.
- Channels can be either **unbuffered** or **buffered**:
  - **Unbuffered**: Sends block until the receiver is ready.
  - **Buffered**: Sends block only when the buffer is full.
- Channels can be directional:
  - `chan T`: Bidirectional (send and receive).
  - `chan<- T`: Send-only.
  - `<-chan T`: Receive-only.

## Examples

### Creating and Using Channels
```go
chan1 := make(chan int) // Unbuffered channel
chan2 := make(chan string, 2) // Buffered channel with capacity 2

// Sending data
chan1 <- 10
chan2 <- "hello"

// Receiving data
val1 := <-chan1
val2 := <-chan2
fmt.Println(val1, val2)
```

### Goroutines and Channels
- Channels are often used to communicate between goroutines.
```go
chan1 := make(chan int)

go func() {
    chan1 <- 10 // Send data to the channel
}()

val := <-chan1 // Receive data from the channel
fmt.Println("Received:", val)
```

### Select Statement
- The `select` statement is used to wait on multiple channel operations.
```go
chan1 := make(chan int)
chan2 := make(chan string)

go func() {
    chan1 <- 10
}()

go func() {
    chan2 <- "pong"
}()

for i := 0; i < 2; i++ {
    select {
    case val := <-chan1:
        fmt.Println("Received from chan1:", val)
    case val := <-chan2:
        fmt.Println("Received from chan2:", val)
    }
}
```

### Directional Channels
- Channels can be restricted to send-only or receive-only.
```go
func sendData(ch chan<- int) {
    ch <- 42
}

func receiveData(ch <-chan int) {
    val := <-ch
    fmt.Println("Received:", val)
}

chan1 := make(chan int)
go sendData(chan1)
go receiveData(chan1)
```

### Buffered Channels
- Buffered channels allow sending multiple values without blocking, up to the buffer capacity.
```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
fmt.Println(<-ch) // Output: 1
fmt.Println(<-ch) // Output: 2
fmt.Println(<-ch) // Output: 3
```

## Additional Notes
- Closing a channel signals that no more data will be sent.
```go
close(chan1)
val, ok := <-chan1 // `ok` is false if the channel is closed
```
- Reading from a closed channel returns the zero value of the channel's type.
- Avoid writing to a closed channel, as it causes a panic.
- Use channels to synchronize goroutines and avoid race conditions.
- The `select` statement with a `default` case can be used for non-blocking operations.
```go
select {
case val := <-ch:
    fmt.Println("Received:", val)
default:
    fmt.Println("No data received")
}
```