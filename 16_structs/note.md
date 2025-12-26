# Notes on Structs in Go

## Key Points
- **Structs** in Go are composite data types that group together fields (variables) under a single name.
- Structs are used to create user-defined types that represent real-world entities.
- Fields in a struct can be of different types, and their names must be unique within the struct.
- Structs can have methods attached to them, allowing you to define behavior for the struct.
- Go does not support inheritance, but you can achieve composition by embedding structs.
- Structs are value types, meaning they are copied when assigned to a new variable or passed to a function.
- Use pointers to structs to modify the original struct.

## Examples

### Declaring and Initializing Structs
- Structs can be declared using the `type` keyword.
```go
type customer struct {
    name  string
    phone string
}

type order struct {
    id        string
    amount    float32
    status    string
    createdAt time.Time // nanosecond precision
    customer
}

order1 := order{
    id:     "1",
    amount: 99.99,
    status: "received",
}
order1.createdAt = time.Now()
fmt.Println(order1)
```

### Accessing and Modifying Struct Fields
- Fields in a struct can be accessed and modified using the dot (`.`) operator.
```go
order1.amount = 120.50
fmt.Println(order1.amount) // Output: 120.50
```

### Struct Methods
- Methods can be attached to structs to define behavior. Use a pointer receiver to modify the struct's fields.
```go
func (o *order) changeStatus(status string) {
    o.status = status
}

func (o order) getAmount() float32 {
    return o.amount
}

order1.changeStatus("shipped")
fmt.Println(order1.status) // Output: shipped
fmt.Println(order1.getAmount()) // Output: 99.99
```

### Struct Composition
- Structs can embed other structs to achieve composition. This allows you to reuse fields and methods.
```go
type customer struct {
    name  string
    phone string
}

type order struct {
    id        string
    amount    float32
    status    string
    createdAt time.Time
    customer
}

order1 := order{
    id:     "1",
    amount: 99.99,
    status: "received",
    customer: customer{
        name:  "John Doe",
        phone: "123-456-7890",
    },
}
fmt.Println(order1.customer.name) // Output: John Doe
```

### Anonymous Structs
- Anonymous structs can be used for quick, one-off data structures without defining a named type.
```go
person := struct {
    name string
    age  int
}{
    name: "Alice",
    age:  30,
}
fmt.Println(person)
```

### Zero Value of Structs
- The zero value of a struct is a struct with all fields set to their zero values.
```go
type product struct {
    name  string
    price float64
}

var p product
fmt.Println(p) // Output: {"" 0}
```

### Using Pointers with Structs
- Use pointers to structs to avoid copying the entire struct and to modify the original struct.
```go
func updateAmount(o *order, newAmount float32) {
    o.amount = newAmount
}

updateAmount(&order1, 150.75)
fmt.Println(order1.amount) // Output: 150.75
```

## Additional Notes
- Structs are the building blocks for creating more complex data structures in Go.
- Use the `new` keyword to create a pointer to a struct.
```go
o := new(order)
o.id = "2"
o.amount = 200.00
fmt.Println(o) // Output: &{2 200 received <zero time> {}}
```
- Structs can be compared using the `==` operator if all their fields are comparable.
- Use tags to add metadata to struct fields, which can be used by libraries like `encoding/json`.
```go
type user struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```