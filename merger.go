package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func copy_file(src, dst string) int64 {
	fmt.Printf("Copying file %v to %v...\n", src, dst)
	source, err := os.Open(
		src)
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


func copy_folder(src, dst string) {
	fmt.Printf("%s -> %s\n", src, dst)
	src_dirnames, src_filenames := list_dir(src)
	for _, dirname := range src_dirnames {
		dirpath := filepath.Join(dst, dirname)
		_, err := os.Stat(dirpath)
		if os.IsNotExist(err) {
			os.MkdirAll(dirpath, os.ModePerm)
		}
	}
	for _, filename := range src_filenames {
		num_bytes := copy_file(
			filepath.Join(src, filename), 
			filepath.Join(dst, filename),
		)
		fmt.Printf("Copied %v bytes\n", num_bytes)
	}
}


func list_dir(wd string) ([]string, []string) {
	fmt.Printf("Listing %s...\n", wd)
	var dirnames, filenames []string
	contents, err := os.ReadDir(wd)
	if err != nil {
		panic(err)
	}
	for _, item := range contents {
		if item.IsDir() {
			dirnames = append(dirnames, item.Name())
		} else {
			filenames = append(filenames, item.Name())
		}
	}
	return dirnames, filenames
}


func main() {
	var dst string
	wd, err := os.Getwd()
	if err != nil {
		panic(err) 
	}
	path_stub := filepath.Join(wd, "/dirs_to_merge")
	dirs_to_merge, _ := list_dir(path_stub)
	sort.Strings(dirs_to_merge)
	for ind, dirname := range dirs_to_merge {
		if ind == 0 {
			dst = filepath.Join(path_stub, dirname)
		} else {
			copy_folder(filepath.Join(path_stub, dirname), dst)
		}
	}
}
