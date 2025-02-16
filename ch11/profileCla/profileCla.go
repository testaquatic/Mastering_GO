package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	cpuFilename := path.Join(os.TempDir(), "cpuProfileCla.out")
	cpuFile, err := os.Create(cpuFilename)
	if err != nil {
		fmt.Println(err)
		return
	}

	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	total := 0
	for i := 2; i < 100_000; i++ {
		n := N1(i)
		if n {
			total += i
		}
	}

	fmt.Println("Total primes:", total)

	total = 0
	for i := 2; i < 100_000; i++ {
		n := N2(i)
		if n {
			total += i
		}
	}
	fmt.Println("Total primes:", total)

	for i := 1; i < 90; i++ {
		n := fibo1(i)
		fmt.Println(n, " ")
	}
	fmt.Println()

	for i := 1; i < 90; i++ {
		n := fibo2(i)
		fmt.Println(n, " ")
	}
	fmt.Println()

	runtime.GC()

	memoryFilename := path.Join(os.TempDir(), "memoryProfileCla.out")
	memory, err := os.Create(memoryFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memory.Close()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50_000_000)
		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep((50 * time.Millisecond))
	}

	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
