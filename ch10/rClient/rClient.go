package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var u1 = User{"admin", "admin"}
var u2 = User{"tsoukalos", "pass"}
var u3 = User{"", "pass"}

const addEndPoint = "/add"
const getEndPoint = "/get"
const deleteEndPoint = "/delete"
const timeEndPoint = "/time"

func deleteEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewBuffer(userMarshall)
	req, err := http.NewRequest(http.MethodDelete, server+deleteEndPoint, u)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resq, err := c.Do(req)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resq == nil {
		return http.StatusNotFound
	}
	defer resq.Body.Close()

	data, err := io.ReadAll(resq.Body)
	fmt.Println("/delete returned: ", string(data))
	if err != nil {
		fmt.Println("Error:", err)
	}

	return resq.StatusCode
}

func getEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewBuffer(userMarshall)

	req, err := http.NewRequest(http.MethodGet, server+getEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resq, err := c.Do(req)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resq == nil {
		return http.StatusNotFound
	}
	defer resq.Body.Close()

	data, err := io.ReadAll(resq.Body)
	fmt.Println("/get returned: ", string(data))
	if err != nil {
		fmt.Println("Error:", err)
	}

	return resq.StatusCode
}

func addEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewBuffer(userMarshall)

	req, err := http.NewRequest(http.MethodPost, server+addEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resq, err := c.Do(req)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resq == nil || (resq.StatusCode == http.StatusNotFound) {
		return resq.StatusCode
	}
	defer resq.Body.Close()

	return resq.StatusCode
}

func timeEndpoint(server string) (int, string) {
	req, err := http.NewRequest(http.MethodPost, server+timeEndPoint, nil)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError, ""
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	return resp.StatusCode, string(data)
}

func slashEndpoint(server, URL string) (int, string) {
	req, err := http.NewRequest(http.MethodPost, server+URL, nil)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError, ""
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp == nil {
		return http.StatusNotFound, ""
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return resp.StatusCode, string(data)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of aguments!")
		fmt.Println("Need: Server")
		return
	}
	server := os.Args[1]

	fmt.Println("/add")
	HTTPcode := addEndpoint(server, u1)
	if HTTPcode != http.StatusOK {
		fmt.Println("u1 Return code:", HTTPcode)
	} else {
		fmt.Println("u1 Data added:", u1, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u2)
	if HTTPcode != http.StatusOK {
		fmt.Println("u2 Return code:", HTTPcode)
	} else {
		fmt.Println("u2 Data added:", u2, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u3)
	if HTTPcode != http.StatusOK {
		fmt.Println("u3 Return code:", HTTPcode)
	} else {
		fmt.Println("u3 Data added:", u3, HTTPcode)
	}

	fmt.Println("/get")
	HTTPcode = getEndpoint(server, u1)
	fmt.Println("/get u1 return code:", HTTPcode)
	HTTPcode = getEndpoint(server, u2)
	fmt.Println("/get u2 return code:", HTTPcode)
	HTTPcode = getEndpoint(server, u3)
	fmt.Println("/get u3 return code:", HTTPcode)

	fmt.Println("/delete")
	HTTPcode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u2)
	fmt.Println("/delete u2 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u3)
	fmt.Println("/delete u3 return code:", HTTPcode)

	fmt.Print("/time")
	HTTPcode, myTime := timeEndpoint(server)
	fmt.Print("/time returned: ", HTTPcode, " ", myTime)
	time.Sleep(time.Second)
	HTTPcode, myTime = timeEndpoint(server)
	fmt.Print("/time returned: ", HTTPcode, " ", myTime)

	fmt.Print("/")
	URL := "/"
	HTTPcode, response := slashEndpoint(server, URL)
	fmt.Print("/ returned: ", HTTPcode, " with response: ", response)

	fmt.Print("/what")
	URL = "/what"
	HTTPcode, response = slashEndpoint(server, URL)
	fmt.Print(URL, " returned: ", HTTPcode, " with response: ", response)
}
