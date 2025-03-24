package main

import (
	"fmt"
	"os"
	"log"
)

func OpenAsmFiles(file string) string{
	data, err := os.ReadFile(file)
	if err != nil{
		log.Fatalf("An error occured while trying to read file %v -> %v", file, err)
	}

	new_data := string(data)
	return new_data
}

func OpenDir(dir string) []string{
	files, err := os.ReadDir(dir)
	if err != nil{
		log.Fatalf("An error occured while trying to read from dir %v -> %v", dir, err)
	}

	// converting the os.DirEntry to string while validating assembly file
	var filename []string
	for _, file := range files{
		name := file.Name()
		if name[len(name)-4:] == ".asm" {filename = append(filename, name)}
	}

	return filename
}

func main() {
	Vars := os.Args[1]
	dir := "./" + Vars
	files := OpenDir(dir)
	
	for _, file := range files{
		content := OpenAsmFiles(dir + "/"+ file)
		fmt.Println("The contents in file ->", file, "->", content)
	}
}
