package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":1234"
var IMAGEPATH = "/tmp/files"

func uploadFile(rw http.ResponseWriter, r *http.Request) {
	filename, ok := mux.Vars(r)["filename"]
	if !ok {
		log.Println("filename value not set!")
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println(filename)
	saveFile(IMAGEPATH+"/"+filename, rw, r)
}

func saveFile(path string, _ http.ResponseWriter, r *http.Request) {
	log.Println("Saving to", path)
	err := saveToFile(path, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func saveToFile(path string, contents io.Reader) error {
	_, err := os.Stat(path)
	if err == nil {
		err = os.Remove(path)
		if err != nil {
			log.Println("Error deleting", path)
			return err
		}
	} else if !os.IsNotExist(err) {
		log.Println("Unexpected error:", err)
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	n, err := io.Copy(f, contents)
	if err != nil {
		return err
	}
	log.Println("Bytes written:", n)

	return nil
}

func createImageDirectory(d string) error {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		log.Println("Creating:", d)
		err = os.MkdirAll(d, 0755)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if err != nil {
		log.Println(err)
		return err
	}

	fileinfo, _ := os.Stat(d)

	mode := fileinfo.Mode()
	if !mode.IsDir() {
		msg := d + " is not a directory!"
		return errors.New(msg)
	}

	return nil
}

func main() {
	err := createImageDirectory(IMAGEPATH)
	if err != nil {
		log.Println(err)
		return
	}

	mux := mux.NewRouter()
	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/files/{filename:[a-zA-Z0-9][a-zA-Z0-9\\.]*[a-zA-Z0-9]}", uploadFile)

	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.Handle("/files/{filename:[a-zA-Z0-9][a-zA-Z0-9\\.]*[a-zA-Z0-9]}", http.StripPrefix("/files/", http.FileServer(http.Dir(IMAGEPATH))))

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	log.Println("Listening to", PORT)

	err = s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
