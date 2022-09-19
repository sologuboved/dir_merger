package main

import (
	"fmt"
	// "io/ioutil"
	// "path/filepath"
	"os"
)


func list(wd string) {
	contents, err := os.ReadDir(wd)
	if err != nil {
		panic(err)
	}
	for _, item := range contents {
		fmt.Println(item.Name(), item.IsDir())
	}
}


func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err) 
	}
	list(wd)
}
