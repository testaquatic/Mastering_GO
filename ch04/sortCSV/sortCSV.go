package main

import (
	"bufio"
	"cmp"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
)

type F1 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type Book1 []F1

var d1 = Book1([]F1{})

type F2 struct {
	Name       string
	Surname    string
	Areacode   string
	Tel        string
	LastAccess string
}

type Book2 []F2

var d2 = Book2([]F2{})

func readFile(filepath string) ([][]string, error) {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	if !fileinfo.Mode().IsRegular() {
		return nil, errors.New("정규 파일이 아닙니다.")
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvR := csv.NewReader(bufio.NewReader(f))

	return csvR.ReadAll()
}

func readCSVFile(filepath string) error {
	firstLine := true
	format1 := true

	lines, err := readFile(filepath)
	if err != nil {
		return err
	}

	for _, line := range lines {
		if firstLine {
			if len(line) == 4 {
				format1 = true
				d1 = slices.Grow(d1, len(lines))
			} else if len(line) == 5 {
				format1 = false
				d2 = slices.Grow(d2, len(lines))
			} else {
				return errors.New("Unknown File Format!")
			}
			firstLine = false
		}

		if format1 {
			if len(line) == 4 {
				temp := F1{
					Name:       line[0],
					Surname:    line[1],
					Tel:        line[2],
					LastAccess: line[3],
				}
				d1 = append(d1, temp)
			}
		} else {
			if len(line) == 5 {
				temp := F2{
					Name:       line[0],
					Surname:    line[1],
					Areacode:   line[2],
					Tel:        line[3],
					LastAccess: line[4],
				}
				d2 = append(d2, temp)
			}
		}
	}

	return nil
}

func sortData(data interface{}) error {
	switch T := data.(type) {
	case Book1:
		d := data.(Book1)
		slices.SortFunc(d, func(a, b F1) int {
			return cmp.Compare(a.Tel, b.Tel)
		})
		list(d)
	case Book2:
		d := data.(Book2)
		slices.SortFunc(d, func(a, b F2) int {
			return cmp.Compare(a.Tel, b.Tel)
		})
		list(d)
	default:
		return fmt.Errorf("Not supported type: %T", T)
	}

	return nil
}

func list(d interface{}) error {
	switch T := d.(type) {
	case Book1:
		data := d.(Book1)
		for _, v := range data {
			fmt.Println(v)
		}
	case Book2:
		data := d.(Book2)
		for _, v := range data {
			fmt.Println(v)
		}
	default:
		return fmt.Errorf("Not supported type: %T", T)
	}

	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "csvfile")
		return
	}

	filepath := flag.Arg(0)

	err := readCSVFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case len(d1) > 0:
		err = sortData(d1)
	case len(d2) > 0:
		err = sortData(d2)
	default:
		fmt.Println("len(d1) == 0 && len(d2) == 0")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
}
