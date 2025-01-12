package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

var timeout = time.Duration(time.Second)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: withDeadline url")
		return
	}
	addr := flag.Arg(0)
	url, err := url.Parse(addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Timeout value:", timeout)

	t := http.Transport{
		Dial: Timeout,
	}
	client := http.Client{
		Transport: &t,
	}

	resp, err := client.Get(url.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
