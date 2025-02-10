package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan<- int, end <-chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			// return
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: select <num>")
		return
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup

	createNumber := make(chan int)
	end := make(chan bool)

	fmt.Printf("Going to create %d random numbers.\n", n)
	wg.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		fmt.Print(<-createNumber, " ")
	}

	end <- true
	wg.Wait()

}
