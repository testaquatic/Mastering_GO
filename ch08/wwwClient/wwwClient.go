package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	c := &http.Client{
		Timeout: 18 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}

	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Status code:", httpData.Status)

	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character set:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("Content length is unknown")
	} else {
		fmt.Println("Content length:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length += n
	}
	fmt.Println("Calculated response data length:", length)
}
