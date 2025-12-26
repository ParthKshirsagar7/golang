```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

```go 
package main
```
In go, every file must belong to a package. The name **main** is special: It tells the Go Compiler that this file should compile as an executable program rather than a shared library.

```go
import "fmt"
```
**import** keyword is an indication to include code from other packages.
"fmt" is a standard library for formatting text.

```go
func main() {}
```
This is the entry point of a go program. On running the compiled binary, the computer looks for the function named main inside main package.

```go
fmt.Println("Hello world")
```
**Println** is a function included in the fmt package which print text to the console and adds a new line at the end.