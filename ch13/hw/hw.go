package main

import "fmt"

func printSlice[T any](s []T) {
	for _, v := range s {
		print(v, " ")
	}
	fmt.Println()
}

func main() {
	printSlice([]int{1, 2, 3})
	printSlice([]string{"a", "b", "c"})
	printSlice([]float64{1.2, -2.33, 4.55})
}
