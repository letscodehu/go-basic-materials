package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	mappings := map[string]string{
		".jpg": "Pictures",
		".png": "Pictures",
		".txt": "Documents",
		".doc": "Documents",
	}
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
	for _, file := range files {
		name := file.Name()
		extension := filepath.Ext(name)
		target, ok := mappings[extension]
		if ok {
			Move(name, target)
		} else {
			fmt.Println("Ignoring", name)
		}
	}
}
func JustPrint(filename string, directory string) {
	fmt.Println("Moving", filename, "to", directory)
}
func Move(filename string, directory string) {
	fmt.Println("Moving", filename, "to", directory)
	os.Mkdir(directory, 0755)
	target := filepath.Join(directory, filename)
	err := os.Rename(filename, target)
	if err != nil {
		fmt.Println(err)
	}
}
