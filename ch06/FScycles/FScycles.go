package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var visited = make(map[string]int)

func walkFunction(path string, info os.FileInfo, err error) error {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return nil
	}
	fileinfo, _ = os.Lstat(path)
	mode := fileinfo.Mode()

	if mode.IsDir() {
		abs, _ := filepath.Abs(path)
		_, ok := visited[abs]
		if ok {
			fmt.Println("Found cycle:", abs)
			return nil
		}
		visited[abs]++
		return nil
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		temp, err := os.Readlink(path)
		if err != nil {
			fmt.Println("os.Readlink():", err)
			return err
		}

		newPath, err := filepath.EvalSymlinks(temp)
		if err != nil {
			return nil
		}

		linkFileInfo, err := os.Stat(newPath)
		if err != nil {
			return err
		}
		linkMode := linkFileInfo.Mode()
		if linkMode.IsDir() {
			fmt.Println("Following...", path, "-->", newPath)

			abs, _ := filepath.Abs(newPath)
			_, ok := visited[abs]
			if ok {
				fmt.Println("Found cycle!", abs)
				return nil
			}
			visited[abs]++
			err = filepath.Walk(newPath, walkFunction)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: ")
		fmt.Printf("FScycles filepath\n")
		return
	}

	path := flag.Arg(0)
	filepath.Walk(path, walkFunction)
}
