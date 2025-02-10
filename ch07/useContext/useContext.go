package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		// 4초
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1() Done:", c1.Err())
		return
		// t초
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1()", r)
	}

	return
}

func f2(t int) {
	c2 := context.Background()
	// t초
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		// 4초
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2() Done:", c2.Err())
		// t초
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}

	return
}

func f3(t int) {
	c3 := context.Background()
	// 2t초
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		// 4초
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3() Done:", c3.Err())
		// t초
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}

	return
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}
