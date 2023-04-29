package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("No target directory specified!")
        fmt.Println("Usage: classifier <directory>")
        return
    }
    directory := os.Args[1]
    fmt.Println(directory)
}
