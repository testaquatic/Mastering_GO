package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func random(r *rand.Rand, min, max int) int {
	return r.Intn(max-min) + min
}

func main() {
	flag.Parse()
	var (
		min   = 0
		max   = 100
		count = 100
		seed  = time.Now().Nanosecond()
	)
	var err error
	if flag.NArg() == 4 {
		if min, err = strconv.Atoi(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		if max, err = strconv.Atoi(flag.Arg(1)); err != nil {
			log.Fatal(err)
		}
		if count, err = strconv.Atoi(flag.Arg(2)); err != nil {
			log.Fatal(err)
		}
		if seed, err = strconv.Atoi(flag.Arg(3)); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Using default values!")
	}

	// `rand.Seed`는 deprecated이다.
	r := rand.New(rand.NewSource(int64(seed)))

	for i := 0; i < count; i++ {
		fmt.Print(random(r, min, max), " ")
	}
	fmt.Println()

}
