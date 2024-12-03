package main

import (
	"embed"
	"fmt"
	"io/fs"
)

func main() {

}

func list(f embed.FS) error {
	return fs.WalkDir(f, ".", walkFunction)
}

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path=%q, isDir=%v\n", path, d.IsDir())
	return nil
}

func extract(f embed.FS, filename string) ([]byte, error) {
	s, err := fs.ReadFile(f, filename)
	if err != nil {
		return nil, err
	}
	return s, err
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
		
	}
}