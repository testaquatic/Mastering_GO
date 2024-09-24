package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()

	var total, nInts, nFloats int
	invalid := make([]string, 0)

	for _, k := range flag.Args() {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, k)
	}

	fmt.Printf("#raad: %d #ints: %d #floats: %d\n", total, nInts, nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, n := range invalid {
			fmt.Println(n)
		}
	}
}
