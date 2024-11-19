package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)

	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[0:n]
}

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Printf("Usage: readSize size filename\n")
		return
	}

	sizeString := flag.Arg(0)
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes := readSize(f, size)
	fmt.Println(string(bytes))
}
