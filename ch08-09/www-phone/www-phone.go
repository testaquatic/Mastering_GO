package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var CSVFILE = "./data.csv"

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int

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

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		index[k.Tel] = i
	}

	return nil
}

func initS(N, S, T string) *Entry {
	if T == "" || S == "" {
		return nil
	}
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)

	return &Entry{
		Name:       N,
		Surname:    S,
		Tel:        T,
		LastAccess: LastAccess,
	}
}

func insert(pS *Entry) error {
	_, ok := index[pS.Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}

	pS.LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	data = append(data, *pS)

	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = slices.Delete(data, i, i+1)
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)

	return &data[i]
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)

	return re.Match(t)
}

func list() string {
	var all string
	for _, k := range data {
		all += k.Name + " " + k.Surname + " " + k.Tel + "\n"
	}

	return all
}

func main() {
	err := readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	fmt.Println(index)

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert/", http.HandlerFunc(insertHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/search", http.HandlerFunc(searchHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
