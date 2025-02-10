package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	PORT := ":8001"
	if flag.NArg() != 0 {
		PORT = ":" + flag.Arg(0)
	}
	fmt.Println("Using port number: ", PORT)

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
