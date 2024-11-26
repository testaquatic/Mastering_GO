package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
)

type XMLrec struct {
	Name    string `xml:"username"`
	Surname string `xml:"surname,omitempty"`
	Year    int    `xml:"creationyear,omitempty"`
}

type JSONrec struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("JSON2XML JSON|XML")
		return
	}

	input := flag.Arg(0)
	fmt.Println(input)

	var jsonData JSONrec
	err := json.Unmarshal([]byte(input), &jsonData)
	if err == nil {
		xmlData := XMLrec{Name: jsonData.Name, Surname: jsonData.Surname, Year: jsonData.Year}
		xmlB, err := xml.Marshal(&xmlData)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(xmlB))
		return
	}

	var xmlData XMLrec
	err = xml.Unmarshal([]byte(input), &xmlData)
	if err != nil {
		fmt.Println("Cannot parse to XML or JSON!")
		return
	}

	jsonData = JSONrec{Name: xmlData.Name, Surname: xmlData.Surname, Year: xmlData.Year}
	jsonB, err := json.Marshal(&jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonB))
}
