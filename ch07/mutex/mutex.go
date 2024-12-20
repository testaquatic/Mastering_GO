package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var m sync.Mutex
var v1 int

func change() {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1 == 10 {
		v1 = 0
		fmt.Print("* ")
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()

	return a
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: mutex.go [goroutine_num]")
		return
	}

	num, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	fmt.Printf("%d ", read())
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			change()
			fmt.Printf("-> %d", read())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("-> %d", read())

}
