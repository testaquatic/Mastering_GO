package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Please provide an agument!")
		return
	}

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	slices.Sort(pathSplit)
	pathSplit = slices.Compact(pathSplit)

	for _, directory := range pathSplit {
		for _, file := range files {
			fullPath := filepath.Join(directory, file)
			// 파일이 존재하는가?
			fileinfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileinfo.Mode()
				// 일반 파일인가?
				if mode.IsRegular() {
					// 실행 파일인가
					if mode&0111 != 0 {
						fmt.Println(fullPath)
					}
				}
			}
		}
	}
}
