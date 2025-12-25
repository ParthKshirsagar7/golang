# Variables

Variables are basically containers which store values.

In go there are three ways to define a variable
1. Standard way:
```go
var name string = "parth"
var a, b, c int = 1, 2, 3;
```
The var keyword is the most explicit way to declare a variable in go.\
Syntax to declare a variable in standard way:\
**var [name] [type] = [value]**

2. Type inference
```go
var country = "India"
```
If you initialize the variable while its declaration, you can skip writing its type explicitly. Go is smart enough to infer data types.\
Just like in the standard way, you cannot change that variable to any other data type after initialization.

3. Shorthand variable declaration (**:=**)
```go
age := 18
```
It is most common way of declaring variables. It's concise and handles declaration and assignment in one go.\
**Note:** := can online be used inside a function. For global variables you will have to use var keyword.