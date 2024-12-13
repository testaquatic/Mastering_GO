package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func createFile(file string) {
	s, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(file), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else if s.Mode().IsDir() {
		log.Fatalf("%s is directory", file)
	}

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(file, "created!")
}

func main() {
	flag.Parse()
	if flag.NArg() != 4 {
		fmt.Println("randomFiles START END FILENAME FILEPATH")
		return
	}
	start, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	end, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := flag.Arg(2)
	path := flag.Arg(3)

	var wg sync.WaitGroup
	for i := start; i <= end; i++ {
		wg.Add(1)
		filepath := fmt.Sprintf("%s/%s%d", path, filename, i)
		go func(f string) {
			defer wg.Done()
			createFile(f)
		}(filepath)
	}
	wg.Wait()
}
