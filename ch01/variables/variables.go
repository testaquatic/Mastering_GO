package main

import (
	"fmt"
	"math"
)

var Global int = 1234
var AnotherGolobal = -5678

func main() {
	var j int
	i := Global + AnotherGolobal
	fmt.Println("Initial j value:", j)
	j = Global
	k := math.Abs(float64(AnotherGolobal))
	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}
