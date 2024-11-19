package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByWord(file string) error {
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

		r := regexp.MustCompile(`[^\s]+`)

		words := r.FindAllString(line, -1)

		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}

	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Usage: byWords files...")
		return
	}
	var err error
	for _, filename := range flag.Args() {
		err = wordByWord(filename)
		if err != nil {
			fmt.Println("error:\n", "\t", err)
		}
	}
}
