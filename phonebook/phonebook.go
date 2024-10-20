package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	// 이름
	Name string
	// 성
	Surname string
	// 전화번호
	Tel        string
	LastAccess string
}

func initS(name, surname, tel string) *Entry {
	return &Entry{
		Name:       name,
		Surname:    surname,
		Tel:        tel,
		LastAccess: fmt.Sprint(time.Now().Unix()),
	}
}

var data = []Entry{}

const CSVFILE = "phonebook.csv"

var index map[string]int

func readCSVFILE() error {
	_, err := os.Stat(CSVFILE)
	if err != nil {
		return err
	}

	f, err := os.Open(CSVFILE)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for line, err := csvReader.Read(); err == nil; line, err = csvReader.Read() {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: strconv.FormatInt(time.Now().Unix(), 10),
		}
		data = append(data, temp)
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func saveCSVFILE(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvWriter := csv.NewWriter(csvfile)
	defer csvWriter.Flush()
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		err = csvWriter.Write(temp)
		if err != nil {
			return err
		}
	}

	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}

	return nil
}

func matchTel(t string) bool {
	r := regexp.MustCompile(`^/d+$`)
	return r.Match([]byte(t))
}

func insert(pS *Entry) error {
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)
	_ = createIndex()

	err := saveCSVFILE(CSVFILE)
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

// 전화번호의 목록을 출력한다.
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = slices.Delete(data, i, i+1)
	delete(index, key)

	err := saveCSVFILE(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}

	err = readCSVFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	fmt.Println("Data has", len(data), "entries.")

	switch flag.Arg(0) {
	case "insert":
		if flag.NArg() != 4 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(flag.Arg(3), "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := initS(flag.Arg(1), flag.Arg(2), t)
		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if flag.NArg() != 2 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(flag.Arg(1), "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if flag.NArg() != 2 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(flag.Arg(1), "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		list()
	// 커맨드를 찾을 수 없는 경우의 응답
	default:
		fmt.Println("Not a valid option")
	}
}
