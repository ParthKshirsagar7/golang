# For

In go, **for** is the only looping construct. 

1. Classic for loop
```go
for i:=0; i<=3; i++ {
    fmt.Println(i)
}
```
This is the most common form consisting of an init statement, a condition and a post statement.

2. While loop equivalent
```go
i:=0
for i<=3 {
    fmt.Println(i)
    i++
}
```

3. Infinite loop
```go
for {
    fmt.Println("hi")
}
```
If you do not profile and condition after for, the loop runs infinitely unit it is broken using the break keyword from inside the for loop.

4. Range
```go
for i := range 3 {
    fmt.Println(i)
}
```
This will print 0 1 and 2 on new lines.

5. break and continue
- **break** keyword is used to break the loop. It is usually used with a condition.
- **continue** keyword skips the current iteration.