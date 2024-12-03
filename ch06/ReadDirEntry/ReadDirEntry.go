package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	contents, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}

	var total int64

	for _, entry := range contents {
		if entry.IsDir() {
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}
			total += temp
		} else {
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}
			total += info.Size()
		}
	}
	return total, nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: ReadDirEntry directory")
		return
	}
	path := flag.Arg(0)
	totalSize, err := GetSize(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Total Size:", totalSize)
}
