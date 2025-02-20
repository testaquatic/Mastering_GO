package examplefunctions

import "fmt"

func ExampleLengthRange() {
	fmt.Println(LengthRange("Mihalis"))
	fmt.Println(LengthRange("Matering Go, 3rd edition!"))
	// output:
	// 7
	// 25
}
