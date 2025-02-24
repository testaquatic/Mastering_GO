package main

import (
	"fmt"
	"reflect"
)

func PrintReflections(s interface{}) {
	fmt.Println("** Reflection")
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}
	fmt.Println()
}

func PrintSlice[T any](s []T) {
	fmt.Println("** Generics")
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.1, 2.2, 3.3})

	PrintReflections([]int{1, 2, 3})
	PrintReflections([]string{"a", "b", "c"})
	PrintReflections([]float64{1.1, 2.2, 3.3})
}
