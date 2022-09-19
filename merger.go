package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
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
	var dirs_to_merge = []string{}
		wd, err := os.Getwd()
	if err != nil {
		panic(err) 
	}
	contents := list(wd + "/dirs_to_merge")
	for _, item := range contents {
		if item.IsDir() == true {
			dirs_to_merge = append(dirs_to_merge, item.Name())
		}
	}
	sort.Strings(dirs_to_merge)
	for ind, filename := range dirs_to_merge {
		if ind != 0 {
			fmt.Println(filename)
		}
	}
}
