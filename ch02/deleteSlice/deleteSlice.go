package main

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("Need an integer value.")
	}

	index := flag.Arg(0)
	i, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Using index", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	if i > len(aSlice)-1 {
		log.Fatalln("Cannot delete element:", i)
	}

	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)

	if i > len(aSlice)-1 {
		log.Fatalln("Cannot delete element:", i)
	}

	aSlice[i] = aSlice[len(aSlice)-1]
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)

	// `slices.Delete()`를 사용해도 될 것 같다.
	// https://pkg.go.dev/slices

	if i > len(aSlice)-2 {
		log.Fatalln("Cannot delete element:", i)
	}
	aSlice = slices.Delete(aSlice, i, i+1)
	fmt.Println("After 3rd deletion:", aSlice)
}
