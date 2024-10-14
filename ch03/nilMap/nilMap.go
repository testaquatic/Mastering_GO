package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil

	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}

	aMap["test"] = 1
	// 프로그램이 충돌한다.
	aMap = nil
	// 경고를 무시한다.
	aMap["test"] = 1
}
