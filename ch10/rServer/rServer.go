package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var user User

var PORT = ":1234"

var DATA = make(map[string]string)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusNotFound)
	Body := "Thanks for visiting!\n"
	_, err := fmt.Fprintf(w, "%s", Body)
	if err != nil {
		log.Println(err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: " + t + "\n"
	_, err := fmt.Fprintf(w, "%s", Body)
	if err != nil {
		log.Println(err)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "%s\n", "Method not allowed!")
		if err != nil {
			log.Println(err)
		}
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	if user.Username != "" {
		DATA[user.Username] = user.Password
		log.Println(DATA)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "%s\n", "Method not allowed!")
		if err != nil {
			log.Println(err)
		}
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll - Error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - Error", http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	_, ok := DATA[user.Username]
	if ok && user.Username != "" {
		log.Println("Found!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", d)
	} else {
		log.Println("Not found!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Map - Resource not found!", http.StatusNotFound)
	}

	return
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodDelete {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "%s\n", "Method not allowed!")
		if err != nil {
			log.Println(err)
		}
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll - Error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - Error", http.StatusBadRequest)
		return
	}
	log.Println(user)

	_, ok := DATA[user.Username]
	if ok && user.Username != "" {
		if user.Password == DATA[user.Username] {
			delete(DATA, user.Username)
			w.WriteHeader(http.StatusOK)
			_, err = fmt.Fprintf(w, "%s\n", d)
			if err != nil {
				log.Println(err)
			}
			log.Println(DATA)
		}
	} else {
		log.Println("User", user.Username, "Not found!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Delete - Rusource not found!", http.StatusNotFound)
	}
	log.Println("After:", DATA)

}

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		PORT = ":" + flag.Arg(0)
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/add", http.HandlerFunc(addHandler))
	mux.Handle("/get", http.HandlerFunc(getHandler))
	mux.Handle("/delete", http.HandlerFunc(deleteHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
