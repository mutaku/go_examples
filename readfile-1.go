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
	filename := os.Args[1]
	return string(filename)
}

func main() {
	fn := getFile()
	contents, _ := loadFile(fn)
	fmt.Println(string(contents.Content))
}
