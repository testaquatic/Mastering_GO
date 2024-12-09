package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%#v\n", wg)

	if len(os.Args) > 1 {
		wg.Add(1)
		wg.Add(1)
	}

	fmt.Println("Going to create 20 goroutines")
	for i := 0; i < 20; i++ {
		go func(num int) {
			wg.Add(1)
			fmt.Printf("%d ", num)
			defer wg.Done()
		}(i)
	}
	wg.Done()
	fmt.Printf("\n%#v\n", wg)

	wg.Wait()
}
