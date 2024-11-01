package main

import (
	"fmt"
	"slices"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grades{
		{"J.", "Lewis", 10},
		{"M.", "Tsoukalos", 7},
		{"D.", "Lewis", 9},
	}

	isSorted := slices.IsSortedFunc(data, func(i, j Grades) int {
		return i.Grade - j.Grade
	})

	if isSorted {
		fmt.Println("It is sorted!")
	} else {
		fmt.Println("It is NOT sorted!")
	}

	slices.SortFunc(data, func(i, j Grades) int {
		return i.Grade - j.Grade
	})

	fmt.Println("By Grade:", data)
}
