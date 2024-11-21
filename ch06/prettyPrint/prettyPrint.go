package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

type KeyVal struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

func generateRandomKeyVal() KeyVal {
	randomChar := func() rune {
		c := rand.Intn(25)
		return rune(int('A') + c)
	}
	var randomKey string
	for i := 0; i < 5; i++ {
		randomKey = randomKey + string(randomChar())
	}

	return KeyVal{Key: randomKey, Val: rand.Intn(100)}
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func JSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	var datas []KeyVal
	for i := 0; i < 2; i++ {
		datas = append(datas, generateRandomKeyVal())
	}

	lastRecord := datas[len(datas)-1]
	fmt.Println("Last record:", lastRecord)
	err := prettyPrint(lastRecord)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := JSONstream(datas)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stream)
}
