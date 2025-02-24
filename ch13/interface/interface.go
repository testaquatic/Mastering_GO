package main

import (
	"fmt"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Print(s interface{}) {
	switch s.(type) {
	case int:
		fmt.Println(s.(int) + 1)
	case float64:
		fmt.Println(s.(float64) + 1.0)
	default:
		println("Unknown data type")
	}
}

func PrintGeneric[T any](s T) {
	fmt.Println(s)
}

func PrintNumeric[T Numeric](s T) {
	fmt.Println(s + 1)
}

func main() {
	Print(12)
	Print(12.0)
	Print("12")

	PrintGeneric(12)	
	PrintGeneric(12.0)
	PrintGeneric("12")

	PrintNumeric(12)
	PrintNumeric(12.0)
}
