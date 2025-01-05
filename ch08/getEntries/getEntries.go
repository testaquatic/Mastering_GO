package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var PORT = ":8765"
var DATAFILE = "/tmp/data.csv"

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type PhoneBook []Entry

var data = PhoneBook{}

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	var tempFileName string

	f, err := os.CreateTemp("", "data*.txt")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tempFileName = f.Name()

	defer os.Remove(tempFileName)

	err = saveCSVFile(tempFileName)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Serving ", tempFileName)

	http.ServeFile(w, r, tempFileName)

	time.Sleep(30 * time.Second)
}

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	defer csvwriter.Flush()
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}

	return nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func main() {
	err := readCSVFile(DATAFILE)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/getContents/", getFileHandler)

	fmt.Println("Starting server on:", PORT)
	err = http.ListenAndServe(PORT, mux)
	fmt.Println(err)
}
