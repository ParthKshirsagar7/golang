# Notes on Packages in Go

## Key Points
- **Packages** in Go are a way to organize and reuse code.
- A package is a collection of Go source files in the same directory that are compiled together.
- The `package` keyword defines the package name at the top of each file.
- Go programs start execution in the `main` package.
- Use the `import` statement to include other packages.
- Packages can be:
  - **Standard library packages** (e.g., `fmt`, `os`).
  - **Third-party packages** (e.g., `github.com/fatih/color`).
  - **Custom packages** (e.g., `auth`, `user`).

## Examples

### Creating and Using Custom Packages
- Define a package by specifying its name at the top of the file.
```go
// File: auth/credentials.go
package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
    fmt.Printf("Logging in the user with username %s and password %s\n", username, password)
}
```

- Use the package in the `main` package by importing it.
```go
// File: main.go
package main

import (
    "github.com/parthkshirsagar7/golang/auth"
)

func main() {
    auth.LoginWithCredentials("admin", "admin")
}
```

### Structs in Custom Packages
- Define structs in a package and export them by capitalizing their names.
```go
// File: user/user.go
package user

type User struct {
    Email string
    Name  string
}
```

- Use the struct in the `main` package.
```go
import "github.com/parthkshirsagar7/golang/user"

func main() {
    user := user.User{
        Email: "1@example.in",
        Name:  "Test User",
    }
    fmt.Println(user.Email, user.Name)
}
```

### Private and Public Identifiers
- Identifiers (functions, variables, structs) that start with an uppercase letter are **exported** (public).
- Identifiers that start with a lowercase letter are **unexported** (private).
```go
// File: auth/session.go
package auth

func extractSession() string { // Private function
    return "Logged in"
}

func GetSession() string { // Public function
    return extractSession()
}
```

### Third-Party Packages
- Use `go get` to install third-party packages.
- Example: The `github.com/fatih/color` package is used for colored terminal output.
```go
import "github.com/fatih/color"

func main() {
    color.Cyan("This is a cyan message")
}
```

### go.mod File
- The `go.mod` file defines the module name and dependencies.
```go
module github.com/parthkshirsagar7/golang

go 1.25.5

require github.com/fatih/color v1.18.0
```
- Use `go mod tidy` to clean up unused dependencies.

## Additional Notes
- Organize code into packages to improve modularity and reusability.
- Use meaningful package names that reflect their functionality.
- Avoid circular dependencies between packages.
- The `init` function in a package is executed automatically when the package is imported.
- Use `go doc` to view documentation for a package.