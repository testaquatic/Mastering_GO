package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

var workers = 4

var sem = semaphore.NewWeighted(int64(workers))

func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)

	return square
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Need #jobs!")
		return
	}

	nJobs, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	var results = make([]int, nJobs)

	ctx := context.TODO()
	for i := range results {
		err = sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("Cannot acquire semaphore:", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			temp := worker(i)
			results[i] = temp
		}(i)
	}

	err = sem.Acquire(ctx, int64(workers))
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
