package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
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
			fmt.Println("error reading file &s", err)
			break
		}
		fmt.Print(line)
	}

	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("파일을 지정하지 않았습니다.")
		return
	}
	for _, fileName := range flag.Args() {
		err := lineByLine(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
