package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	var k *aStructure
	fmt.Println(k)
	// 아래의 경고는 무시한다.
	if k == nil {
		k = new(aStructure)
	}

	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nill")
	}
}
