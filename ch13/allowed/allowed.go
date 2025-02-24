package main

import "fmt"

func Same[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println("4 = 3 is", Same(4, 3))
	fmt.Println("aa = aa is", Same("aa", "aa"))
	fmt.Println("4.1 = 4.15 is", Same(4.1, 4.15))
	// 아래 코드는 오류가 발생한다.
	// fmt.Println("[]int{1,2} = []int{1,3}", Same([]int{1,2}, []int{1,3}))
}
