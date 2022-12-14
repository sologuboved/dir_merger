package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func TimeIt(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s %s\n", name, fmt.Sprintf("took %s", elapsed))
}


func CopyFile(src, dst string) int64 {
	fmt.Printf("\nCopying file %v to %v...\n", src, dst)
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


func MergeDirs(src, dst string) {
	fmt.Printf("%s -> %s\n", src, dst)
	src_dirnames, src_filenames := ListDir(src)
	for _, filename := range src_filenames {
		num_bytes := CopyFile(
			filepath.Join(src, filename), 
			filepath.Join(dst, filename),
		)
		fmt.Printf("Copied %v bytes\n", num_bytes)
	}
	for _, dirname := range src_dirnames {
		dst_dirpath := filepath.Join(dst, dirname)
		_, err := os.Stat(dst_dirpath)
		if os.IsNotExist(err) {
			os.MkdirAll(dst_dirpath, os.ModePerm)
		}
		MergeDirs(filepath.Join(src, dirname), dst_dirpath)
	}
}


func ListDir(wd string) ([]string, []string) {
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
	defer TimeIt(time.Now(), "main")
	var dst string
	wd, err := os.Getwd()
	if err != nil {
		panic(err) 
	}
	path_stub := filepath.Join(wd, "/dirs_to_merge")
	dirs_to_merge, _ := ListDir(path_stub)
	sort.Strings(dirs_to_merge)
	for ind, dirname := range dirs_to_merge {
		if ind == 0 {
			dst = filepath.Join(path_stub, dirname)
		} else {
			MergeDirs(filepath.Join(path_stub, dirname), dst)
		}
	}
}
