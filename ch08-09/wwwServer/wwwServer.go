package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	flag.Parse()
	PORT := ":8001"
	if flag.NArg() != 0 {
		PORT = ":" + flag.Arg(0)
	}
	fmt.Println("Using port number: ", PORT)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
