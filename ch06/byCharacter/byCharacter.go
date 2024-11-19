package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}

	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Usage: byCharacter filenames...")
		return
	}
	for _, filename := range flag.Args() {
		err := charByChar(filename)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	}
}
