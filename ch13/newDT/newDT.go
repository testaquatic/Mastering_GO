package main

import (
	"errors"
	"fmt"
)

type TreeLast[T any] []T

func (t TreeLast[T]) replaceLast(element T) (TreeLast[T], error) {
	if len(t) == 0 {
		return nil, errors.New("tree is empty")
	}
	t[len(t)-1] = element
	return t, nil
}

func main() {
	tempStr := TreeLast[string]{"aa", "bb"}
	fmt.Println(tempStr)
	tempStr.replaceLast("cc")
	fmt.Println(tempStr)

	tempInt := TreeLast[int]{1, 2}
	fmt.Println(tempInt)
	tempInt.replaceLast(3)
	fmt.Println(tempInt)
}
