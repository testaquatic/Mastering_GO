package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	flag.Parse()
	value := int64(8)
	if flag.NArg() != 1 {
		fmt.Println("Using default values!")
	} else {
		temp, err := strconv.ParseInt(flag.Arg(0), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		value = temp
	}

	byteLen := (value/3 + 1) * 4
	bytes, err := generateBytes(int64(byteLen))
	if err != nil {
		log.Fatal(err)
	}

	random := base64.URLEncoding.EncodeToString(bytes)
	fmt.Println(random[:value])
}
