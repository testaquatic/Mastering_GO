package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: varGoroutines COUNT")
		return
	}
	countString := flag.Arg(0)
	count, err := strconv.Atoi(countString)
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
