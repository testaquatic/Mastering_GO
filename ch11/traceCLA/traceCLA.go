package main

import (
	"fmt"
	"os"
	"path"
	"runtime/trace"
	"time"
)

func main() {
	filename := path.Join(os.TempDir(), "traceCLA.out")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	for i := 0; i < 3; i++ {
		s := make([]byte, 50_000_000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}

	for i := 0; i < 5; i++ {
		s := make([]byte, 100_000_000)
		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep(time.Millisecond)
	}
}
