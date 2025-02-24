package main

import "fmt"

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("4 + 3 =", Add(4, 3))
	fmt.Println("4.5 + 3.5 =", Add(4.5, 3.5))
	// 아래 코드는 컴파일 오류가 발생한다.
	// fmt.Println("4 + 3.5 =", Add(int(4), 3.5))
}
