package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":1234"

func timeHandler(w http.ResponseWriter, r *http.Request) {}

func addHandler(w http.ResponseWriter, r *http.Request) {}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s from %s", r.RequestURI, r.Host)
		next.ServeHTTP(w, r)
	})
}

func anotherMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(PORT) != 0 {
			log.Printf("Using port: %s", PORT)
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if flag.NArg() >= 1 {
		PORT = flag.Arg(0)
	}

	mux := mux.NewRouter()
	mux.Use(middleware)

	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/time", timeHandler)

	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/add", addHandler)
	getMux.Use(anotherMiddleWare)

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
