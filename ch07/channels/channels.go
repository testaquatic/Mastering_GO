package main

import (
	"fmt"
)

func writeToChannel(c chan<- int, x int) {
	c <- x
	close(c)
}

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main() {
	var ch chan bool = make(chan bool)

	go printer(ch, 5)

	for val := range ch {
		fmt.Println(val, " ")
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Println(<-ch, " ")
	}
	fmt.Println()
}
