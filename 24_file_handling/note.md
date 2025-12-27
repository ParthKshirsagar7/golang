# Notes on File Handling in Go

## Key Points
- The `os` package in Go provides functions for file handling, including creating, opening, reading, writing, and deleting files.
- File handling involves working with file descriptors and managing resources properly (e.g., using `defer` to close files).
- Go supports both buffered and unbuffered file operations.
- The `bufio` package provides utilities for buffered reading and writing, which are useful for large files.

## Examples

### Opening a File
- Use `os.Open` to open a file in read-only mode.
```go
f, err := os.Open("example.txt")
if err != nil {
    panic(err) // Handle errors appropriately
}
defer f.Close() // Ensure the file is closed when done
```

### Getting File Information
- Use the `Stat` method to retrieve metadata about a file.
```go
fileInfo, err := f.Stat()
if err != nil {
    panic(err)
}

fmt.Println("Is Directory:", fileInfo.IsDir())
fmt.Println("Filename:", fileInfo.Name())
fmt.Println("Size:", fileInfo.Size())
fmt.Println("Permissions:", fileInfo.Mode())
fmt.Println("Last Modified:", fileInfo.ModTime())
```

### Reading a File
- Read the contents of a file using `Read`.
```go
buf := make([]byte, fileInfo.Size())
_, err = f.Read(buf)
if err != nil {
    panic(err)
}
fmt.Println("Data:", string(buf))
```

- Alternatively, use `os.ReadFile` to load the entire file into memory (suitable for small files).
```go
data, err := os.ReadFile("example.txt")
if err != nil {
    panic(err)
}
fmt.Println("Data:", string(data))
```

### Listing Files in a Directory
- Use `ReadDir` to list files in a directory.
```go
dir, err := os.Open("../")
if err != nil {
    panic(err)
}
defer dir.Close()

info, err := dir.ReadDir(-1)
if err != nil {
    panic(err)
}

for _, file := range info {
    fmt.Println(file.Name())
}
```

### Writing to a File
- Use `os.Create` to create a new file (or truncate an existing file) and write data to it.
```go
file, err := os.Create("example2.txt")
if err != nil {
    panic(err)
}
defer file.Close()

_, err = file.Write([]byte("Hi golang"))
if err != nil {
    panic(err)
}
```

### Copying Data Between Files
- Use `bufio` for buffered reading and writing to copy data between files.
```go
sourceFile, err := os.Open("example.txt")
if err != nil {
    panic(err)
}
defer sourceFile.Close()

destFile, err := os.Create("example2.txt")
if err != nil {
    panic(err)
}
defer destFile.Close()

reader := bufio.NewReader(sourceFile)
writer := bufio.NewWriter(destFile)

for {
    b, err := reader.ReadByte()
    if err != nil {
        if err.Error() != "EOF" {
            panic(err)
        }
        break
    }

    e := writer.WriteByte(b)
    if e != nil {
        panic(e)
    }
}

writer.Flush()
fmt.Println("Written to new file successfully")
```

### Deleting a File
- Use `os.Remove` to delete a file.
```go
err := os.Remove("example2.txt")
if err != nil {
    panic(err)
}
fmt.Println("File deleted successfully")
```

## Additional Notes
- Always handle errors when working with files, as file operations can fail for various reasons (e.g., missing files, permission issues).
- Use `defer` to ensure files are closed properly, even if an error occurs.
- For large files, avoid loading the entire file into memory; use buffered reading instead.
- The `io` and `bufio` packages provide additional utilities for file handling, such as buffered reading and writing.