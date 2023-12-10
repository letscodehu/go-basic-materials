package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	file, _ := os.Open("numbers.txt")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Closing file due to panic:", r)
		} else {
			fmt.Println("Closing file normally")
		}
		file.Close()
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		fmt.Println(500 / number)
	}
}
func main2() {
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
			err := Move(name, target)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Ignoring", name)
		}
	}
}

func Move(filename string, directory string) error {
	fmt.Println("Moving", filename, "to", directory)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Println("Creating directory:", directory)
		if err := os.Mkdir(directory, 0755); err != nil {
			return fmt.Errorf("Error creating directory: %s", err)
		}
	}
	target := filepath.Join(directory, filename)
	if err := os.Rename(filename, target); err != nil {
		return fmt.Errorf("Error moving file: %s", err)
	}
	fmt.Println("Move successful.")
	return nil
}
