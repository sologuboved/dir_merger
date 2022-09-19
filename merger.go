package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func copy_file(src, dst string) int64 {
	source, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer destination.Close()

	num_bytes, err := io.Copy(destination, source)
	if err != nil {
		panic(err)
	}

	return num_bytes
}


func list(wd string) []fs.DirEntry {
	fmt.Printf("Listing %s...\n", wd)
	contents, err := os.ReadDir(wd)
	if err != nil {
		panic(err)
	}
	return contents
}


func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err) 
	}
	contents := list(wd + "/dirs_to_merge")
	for _, item := range contents {
		fmt.Println(item.Name(), item.IsDir())
	}
}
