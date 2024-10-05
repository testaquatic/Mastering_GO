package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])

	fmt.Println(aSlice[l-2 : l])
	fmt.Println(aSlice[l-2:])

	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))
	t = aSlice[2:4:10]
	fmt.Println(len(t), cap(t))

	t = aSlice[0:5:6]
	fmt.Println(len(t), cap(t))
}
