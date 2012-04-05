//Read in a file provided as argument and print out the contents
package main

import (
	"io/ioutil"
	"flag"
	"fmt"
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
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Need a file argument.")
	} else { 
		filename = flag.Arg(0)
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
