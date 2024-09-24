package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	arguments := flag.Args()
	if len(arguments) == 0 {
		fmt.Println("Please provide an agument!")
		return
	}
	file := arguments[0]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
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
					return
				}
			}
		}
	}
}
