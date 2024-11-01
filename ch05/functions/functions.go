package main

import "fmt"

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func sortTwo(x, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func main() {
	n := 10
	d, s := doubleSquare(n)

	fmt.Println("Double of", n, "is", d)
	fmt.Println("Square of", n, "is", s)

	anF := func(param int) int {
		return param * param
	}

	fmt.Println("anF of", n, "is", anF(n))

	fmt.Println(sortTwo(1, -3))
	fmt.Println(sortTwo(-1, 0))
}
