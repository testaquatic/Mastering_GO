package main

<<<<<<< HEAD
import (
	"fmt"
	"math/rand"

	"github.com/testaquatic/post05"
)

var MIN = 0
var MAX = 25

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}

	return temp
}

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5432
	post05.Username = "aquatic"
	post05.Password = "pass"
	post05.Database = "go"

	data, err := post05.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	randomUsername := getString(5)

	t := post05.Userdata{
		Username:    randomUsername,
		Name:        "Aquatic",
		Surname:     "Life",
		Description: "Opabinia",
	}

	id := post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}
	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = post05.Userdata{
		Username:    randomUsername,
		Name:        "Aquatic",
		Surname:     "Life",
		Description: "Wiwaxia",
	}

	err = post05.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}
=======
var MIN = 0
var M

func main() {
	
}
>>>>>>> e5392f9 (﻿post05-2)
