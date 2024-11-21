package main

import (
	"encoding/json"
	"fmt"
)

type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
}

type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main() {
	noEmptyVar := NoEmpty{
		Name: "Nihalis",
	}
	noEmptyVarJson, err := json.Marshal(&noEmptyVar)
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}
	fmt.Printf("noEmptyVar decoded with value %s\n", string(noEmptyVarJson))

	password := Password{
		Name: "Mihalis",
		Pass: "myPassword",
	}
	passwordJson, err := json.Marshal(&password)
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}
	fmt.Printf("password decoded with value %s\n", string(passwordJson))
}
