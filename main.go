package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No target directory specified!")
		fmt.Println("Usage: classifier <directory>")
		os.Exit(1)
	}
	directory := os.Args[1]
	info, err := os.Stat(directory)
	if err != nil {
		fmt.Println(directory, "does not exists.")
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Println(directory, "is not a directory.")
		os.Exit(1)
	}
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err, "is not a directory.")
		os.Exit(1)
	}
	for index, file := range files {
		fmt.Println(index, file.Name())
	}
}
