package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

func main() {
	readdirDemo()
	walkDemo()
}

func readdirDemo() {
	// dir is *os.File (pointer to File struct)
	dir, err := os.Open(".")
	fmt.Println(reflect.TypeOf(dir))

	if err != nil {
		fmt.Println("Open: could not open directory")
		return
	} else {
		defer dir.Close()
	}

	// ReadDir reads some number of entries, we pass -1 to read all
	// info is a slice of *os.fileStat (pointer to fileStat struct)
	info, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir: could nor read directory")
		return
	}

	for _, obj := range info {
		fmt.Println(obj, reflect.TypeOf(obj), obj.Name())
	}
}

func walkDemo() {
	// a function is passed to Walk() which gets called for each recursive entry
	// it takes the path, a FileInfo struct (os.Stat()) and an error
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
