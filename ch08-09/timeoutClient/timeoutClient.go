package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var myUrl string
var delay int = 5
var wg sync.WaitGroup

type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer wg.Done()

	data := make(chan myData, 1)
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	req, _ := http.NewRequestWithContext(c, http.MethodGet, myUrl, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	case <-c.Done():
		fmt.Println("The request was canceled!")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server Response: %s\n", realHTTPData)
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: timeoutClient <url> <delay>")
		return
	}

	myUrl = flag.Arg(0)
	if flag.NArg() == 2 {
		t, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay:", delay)
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s\n", myUrl)
	wg.Add(1)
	go connect(c)
	wg.Wait()
	fmt.Println("Exiting...")
}
