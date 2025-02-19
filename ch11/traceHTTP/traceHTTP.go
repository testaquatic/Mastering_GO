package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptrace"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: URL")
		return
	}
	URL := flag.Arg(0)
	client := http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)

	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(gci httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", gci)
		},
		DNSDone: func(di httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", di)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requesting data from server!")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
}
