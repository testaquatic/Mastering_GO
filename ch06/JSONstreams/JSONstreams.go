package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Deserialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

type KeyVal struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

func main() {
	kvSlice := []KeyVal{
		{Key: "XVLBZ", Val: 16},
		{Key: "BAICM", Val: 89},
	}

	buffer := new(bytes.Buffer)
	jsonEn := json.NewEncoder(buffer)
	err := Serialize(jsonEn, kvSlice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("After Serialize:%s", buffer.String())

	deserialized := []KeyVal{}
	jsonDe := json.NewDecoder(buffer)
	err = Deserialize(jsonDe, &deserialized)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("After Deserialize:")
	for i, kv := range deserialized {
		fmt.Println(i, kv)
	}
}
