package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("folder?:", fileInfo.IsDir())
	// fmt.Println("filename:", fileInfo.Name())
	// fmt.Println("size:", fileInfo.Size())
	// fmt.Println("file permision:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// read file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// buf := make([]byte, fileInfo.Size())

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data:", string(buf), d)

	// os.ReadFile() loads the file into the memory at once. This is fine for small files but would create a problem for larger files
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// info, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range info {
	// 	fmt.Println(file.Name())
	// }

	// create a file
	// file, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // file.WriteString("This is being written through code")
	// // file.WriteString("Testing overwritting or appending")

	// data := []byte("Hi golang")
	// file.Write(data)

	// read and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Written to new file successfully")

	// delete a file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully")
}
