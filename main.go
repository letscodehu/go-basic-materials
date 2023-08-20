package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	// n := 0
	// for n < len(files) {
	// 	fmt.Println(n, files[n].Name())
	// 	n++
	// }
	// for i := 0; i < len(files); i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i, files[i].Name())
	// }
	for index, file := range files {
		name := file.Name()
		extension := filepath.Ext(name)
		fmt.Println(index, name, extension)
	}
	// for {
	// 	fmt.Println("Test")
	// }
}
