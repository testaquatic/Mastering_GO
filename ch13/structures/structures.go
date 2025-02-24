package main

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
}

type list[T any] struct {
	head *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{data: data, next: nil}

	if l.head == nil {
		l.head = &n
		return
	}

	if l.head.next == nil {
		l.head.next = &n
		return
	}

	temp := l.head
	l.head = l.head.next
	l.add(data)
	l.head = temp
}

func main() {
	var myList list[int]

	fmt.Println(myList)

	myList.add(12)
	myList.add(13)
	myList.add(14)
	myList.add(13)

	for {
		fmt.Println("*", myList.head)
		if myList.head == nil {
			break
		}
		myList.head = myList.head.next
	}
}
