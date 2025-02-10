package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:

		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Please give an integer!")
		return
	}

	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random numbers.\n", n)
	go monitor()

	var wg sync.WaitGroup

	for r := 0; r < n; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
