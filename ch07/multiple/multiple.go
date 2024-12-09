package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: multiple COUNT")
		return
	}
	countString := flag.Arg(0)
	count, err := strconv.Atoi(countString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
