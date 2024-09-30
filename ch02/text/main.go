package main

import (
	"fmt"
)

func main() {
	aString := "Hello World! 유니코드"
	fmt.Println("First character", string(aString[0]))

	r := '룬'
	fmt.Println("As an int32 value:", r)
	// 아래 코드의 경고는 무시할 것
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	for _, v := range aString {
		fmt.Printf("0x%x ", v)
	}
	fmt.Println()
}
