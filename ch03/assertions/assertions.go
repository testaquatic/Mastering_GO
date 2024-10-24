package main

import "fmt"

func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()
	number := anInt.(int)
	number++
	fmt.Println(number)

	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	i := anInt.(int)
	fmt.Println("i:", i)
	_ = anInt.(bool)
}
