package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed static/image.png
var f1 []byte

//go:embed static/textfile
var f2 string

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

	return err
}
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Print select 1|2")
		return
	}

	fmt.Println("f1:", len(f1), "f2", len(f2))

	switch flag.Arg(0) {
	case "1":
		filename := "/tmp/temporary.png"
		err := writeToFile(f1, filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":
		fmt.Print(f2)
	default:
		fmt.Println("Not a valid option!")
	}
}
