package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Panic(err)
	}

	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %v of type %T\n", input, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %v of type %T\n", input, input)
	input = string(n)
	fmt.Printf("string() %v of type %T\n", input, input)
}
