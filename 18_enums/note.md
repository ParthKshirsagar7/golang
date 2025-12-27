# Notes on Enums in Go

## Key Points
- Go does not have a built-in `enum` type like some other languages (e.g., C# or Java).
- Enums in Go are typically implemented using `const` blocks with the `iota` identifier.
- `iota` is a special identifier that simplifies the creation of incrementing constants.
- Enums are often associated with a custom type (e.g., `int`, `string`) to provide better type safety and readability.

## Examples

### Declaring Enums with `iota`
- Use a `const` block and `iota` to define enumerated constants.
```go
type OrderStatus int

const (
    Received OrderStatus = iota
    Confirmed
    Prepared
    Delivered
)
```
- The `iota` identifier starts at `0` and increments by `1` for each constant in the block.

### Using Enums
- Enums can be used as function parameters or fields in structs.
```go
func changeOrderStatus(status OrderStatus) {
    fmt.Println("Changing order status to", status)
}

func main() {
    changeOrderStatus(Delivered) // Output: Changing order status to 3
}
```

### Adding String Representation
- Implement the `String()` method to provide a human-readable representation of enum values.
```go
func (status OrderStatus) String() string {
    switch status {
    case Received:
        return "Received"
    case Confirmed:
        return "Confirmed"
    case Prepared:
        return "Prepared"
    case Delivered:
        return "Delivered"
    default:
        return "Unknown"
    }
}

func main() {
    fmt.Println(Delivered) // Output: Delivered
}
```

### Custom Enum Types
- Enums can use other types, such as `string`, for better readability.
```go
type PaymentStatus string

const (
    Pending   PaymentStatus = "Pending"
    Completed PaymentStatus = "Completed"
    Failed    PaymentStatus = "Failed"
)

func main() {
    var status PaymentStatus = Completed
    fmt.Println(status) // Output: Completed
}
```

## Additional Notes
- Enums in Go are flexible and can be extended with methods to add functionality.
- Use enums to represent a fixed set of related constants, such as states, statuses, or categories.
- The `iota` identifier resets to `0` in each `const` block.
- Enums can be combined with `switch` statements for better control flow.
```go
switch status {
case Received:
    fmt.Println("Order received")
case Delivered:
    fmt.Println("Order delivered")
}
```