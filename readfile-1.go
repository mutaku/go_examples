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

func getFile() (filename string) {
	if len(os.Args) < 2 {
		fmt.Println("Need a file argument.")
	} else { 
		filename = os.Args[1]
	}
	return string(filename)
}

func main() {
	fn := getFile()
	if fn != "" {
		contents, _ := loadFile(fn)
		fmt.Println(contents.Name, ":\n", string(contents.Content))
	}
}
