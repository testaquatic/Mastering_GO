package main

import (
	"flag"
	"fmt"
	"strconv"
)

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x

		return min, max
	}

	min = x
	max = y
	return
}

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("flag.NArg() != 2")
		return
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println(err)
		return
	}

	x, y = minMax(x, y)
	fmt.Println(x, y)
	fmt.Println(minMax(x, y))
}
