package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var result = make(chan bool)

func timeout(t time.Duration) {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
	}()

	select {
	case <-temp:
		result <- false
	case <-time.After(t):
		result <- true
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("timeOut2 time(ms)")
		return
	}

	timeMs, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Millisecond * time.Duration(timeMs)
	fmt.Println("Timeout period is", t)

	go timeout(time.Duration(timeMs) * time.Millisecond)

	if <-result {
		fmt.Println("Time out!")
	} else {
		fmt.Println("OK")
	}
}
