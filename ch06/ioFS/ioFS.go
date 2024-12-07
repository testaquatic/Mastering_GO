package main

import (
	"embed"
	"fmt"
	"io/fs"
<<<<<<< HEAD
	"os"
)

//go:embed static
var f embed.FS

var searchString string

func main() {
	err := list(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	searchString = "file.txt"
	err = search(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer, err := extract(f, "static/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeToFile(buffer, "/tmp/IOFS.txt")
	if err != nil {
		fmt.Println(err)
	}
=======
)

func main() {

>>>>>>> Mastering_GO/main
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

<<<<<<< HEAD
func search(f embed.FS) error {
	return fs.WalkDir(f, ".", walkSearch)
}

=======
>>>>>>> Mastering_GO/main
func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
<<<<<<< HEAD
		fileInfo, err := fs.Stat(f, path)
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size())
		return nil
	}

	return nil
}

func writeToFile(s []byte, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	n, err := fd.Write(s)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}
=======
		
	}
}
>>>>>>> Mastering_GO/main
