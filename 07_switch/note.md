# Switch Statement in Go

## Overview
A switch statement is used to execute different code blocks based on different conditions. It's an alternative to multiple if-else statements and makes code cleaner and more readable.

## Basic Switch Statement

### Syntax
```go
switch expression {
case value1:
    // code executes if expression == value1
case value2:
    // code executes if expression == value2
default:
    // code executes if expression doesn't match any case
}
```

### Example
```go
i := 5
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
default:
    fmt.Println("other")
}
// Output: other
```

**Key Points:**
- Each case statement tests if the switch expression equals that case value
- If no case matches, the `default` case executes (optional)
- Only one case block executes (no automatic fallthrough unless specified)

---

## Multiple Conditions in a Single Case

### Syntax
You can test multiple values in a single case by separating them with commas.

### Example
```go
switch time.Now().Weekday() {
case time.Sunday, time.Saturday:
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

**Key Points:**
- Multiple values can be checked in a single case separated by commas
- If the expression matches any of the values, that case executes
- This reduces code duplication when multiple values have the same logic

---

## Type Switch

### Overview
A type switch compares types instead of values. It's used to discover the dynamic type of an interface{} value.

### Syntax
```go
switch v := x.(type) {
case type1:
    // v is of type1
case type2:
    // v is of type2
default:
    // v is some other type
}
```

### Example
```go
var k interface{} = false
switch i := k.(type) {
case int:
    fmt.Println("Its an integer")
case string:
    fmt.Println("Its a string")
case bool:
    fmt.Println("Its a boolean")
default:
    fmt.Println("Other", i)
}
// Output: Its a boolean
```

**Key Points:**
- Type switches use the syntax `x.(type)` where x is an interface{} value
- The variable `i` (in the example) gets assigned the value of `k` converted to the matched type
- In the `default` case, `i` retains its original type (interface{})
- Commonly used to handle interface{} values of unknown types

---

## Fallthrough Statement

### Overview
By default, once a case matches, other cases are not executed. The `fallthrough` keyword allows execution to continue to the next case.

### Syntax
```go
switch expression {
case value1:
    // code
    fallthrough  // execution continues to the next case
case value2:
    // code also executes if value1 matched
}
```

### Example
```go
switch num := 2; num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}
// Output: Two
//         Three
```

**Key Points:**
- `fallthrough` forces execution to continue to the next case without checking its condition
- The next case's code executes regardless of whether its condition matches
- Useful when you want to execute code from multiple cases for a single match
- **Note:** fallthrough goes to the next case in order, not to a specific case

---

## Summary Table

| Feature | Description |
|---------|-------------|
| **Basic Switch** | Compares a value against multiple cases |
| **Default Case** | Executes when no case matches (optional) |
| **Multiple Values** | Test multiple values in one case (comma-separated) |
| **Type Switch** | Compares types of interface{} values |
| **Fallthrough** | Continues execution to the next case |

---

## Advantages of Switch Statements
1. More readable than multiple if-else statements
2. Better performance for many conditions
3. Type switches provide elegant type checking for interfaces
4. Cleaner syntax for testing single expressions against multiple values