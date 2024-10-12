package main

// https://pkg.go.dev/sort 이 문서를 읽어보면 일반적으로 slices.sort를 권장하는 것 같다.
// slices.Sort를 사용해서 코드를 작성한다.

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	sInts := []int{1, 0, 2, -3, 4, -20}
	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	sStrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}

	fmt.Println("sInts original:", sInts)
	slices.Sort(sInts)
	fmt.Println("sInts:", sInts)
	slices.SortFunc(sInts, func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Println("Reverse:", sInts)

	fmt.Println("sFloats original:", sFloats)
	slices.Sort(sFloats)
	fmt.Println("sFloats:", sFloats)
	slices.SortFunc(sFloats, func(a, b float64) int { return cmp.Compare(b, a) })
	fmt.Println("Reverse:", sFloats)

	fmt.Println("sStrings original:", sStrings)
	slices.Sort(sStrings)
	fmt.Println("sStrings:", sStrings)
	slices.SortFunc(sStrings, func(a, b string) int { return cmp.Compare(b, a) })
	fmt.Println("Reverse", sStrings)
}
