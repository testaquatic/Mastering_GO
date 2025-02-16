package main

import (
	"math"
	"time"
)

func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)

	return int64(fibo2(n-1)) + int64(fibo2(n-2))
}

func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Microsecond)

	return fn[n]
}

func N1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}

	return true
}

func N2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}

	return true
}
