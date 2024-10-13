package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

const (
	MIN = 0
	MAX = 94
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {
	flag.Parse()
	defaultLen := int64(8)
	if flag.NArg() == 1 {
		var err error
		defaultLen, err = strconv.ParseInt(flag.Arg(0), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Using default values...")
	}
	fmt.Println(getString(defaultLen))
}
