package main

import (
	"fmt"
	"runtime"
	"sync"
)

var m sync.Mutex
var w sync.WaitGroup

func function() {
	m.Lock()
	fmt.Println("Locked!")
}

func main() {
	count := runtime.GOMAXPROCS(0)

	for i := 0; i < count; i++ {
		w.Add(1)
		go func() {
			function()
			w.Done()
		}()
	}
	w.Wait()
}
