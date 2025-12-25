# Golang notes

## Reasons to choose Golang
1. **Build time**: Golang has very low build time which optimizes resources used in a DevOps cycle. Lower build time also enable lesser testing and deployment times.
2. **Fast startup**: After deplying and running a go program, the time it takes to be up and running is very less. This can help in scaling of applications as in this use case lesser time taken by the program to run will lead to lower downtime and ultimately a better user experience.
3. **Performance & Efficiency**: Go programs use system resources efficiently.
4. **Concurrency model**: Go has Goroutines which are lightweight threads compared to typical OS threads. One can run hundreds of thousands of goroutines on a standard laptop. It basically enables you to run multiple operations without exhausting the memory.
5. **Static typing and compilation**: Go is a statically typed language. It might feel like it is not due to the **walrus operator :=** but during compile time, the type of variable assigned using walrus operator is fixed Go throws an error if anything such happens.