package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFILE(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("csvData input output")
		return
	}

	input := flag.Arg(0)
	output := flag.Arg(1)
	lines, err := readCSVFILE(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		fmt.Print(temp)
	}

	err = saveCSVFile(output)
	if err != nil {
		fmt.Println(err)
		return
	}
}
