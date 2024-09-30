package main

import (
	"fmt"
	"math"
)

type Digit int
type Power2 int

const PI = math.Pi

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		P2_0 Power2 = 1 << iota
		_
		P2_2
		_
		P2_4
		_
		P2_6
	)

	fmt.Println("2^0", P2_0)
	fmt.Println("2^2", P2_2)
	fmt.Println("2^4", P2_4)
	fmt.Println("2^6", P2_6)
}
