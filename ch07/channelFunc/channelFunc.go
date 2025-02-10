package main

import "fmt"

func printer(ch chan<- bool) {
	ch <- true
}

func writeToChannel(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
}

func main() {
	boolCh := make(chan bool)
	go printer(boolCh)
	fmt.Println(<-boolCh)

	inCh := make(chan int, 1)
	outCh := make(chan int, 1)

	writeToChannel(outCh, 1)
	f2(outCh, inCh)
	fmt.Println(<-inCh)

}
