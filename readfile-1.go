//Read in a file provided as argument and print out the contents
package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

type File struct {
	Name string
	Content []byte
}

func loadFile(name string) (*File, error) {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return &File{Name: name, Content: content}, nil
}

func getFile() string {
	var filename string
	// so many returns here. This is wrong... but... works.
	if len(os.Args) < 2 {
		fmt.Println("Need a file argument.")
		filename := ""
		return filename
	} else { 
		filename := os.Args[1]
		return filename
	}
	return string(filename)
}

func main() {
	fn := getFile()
	if fn != "" {
		contents, _ := loadFile(fn)
		fmt.Println(string(contents.Content))
	}
}
