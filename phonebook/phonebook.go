package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Entry struct {
	// 이름
	Name string
	// 성
	Surname string
	// 전화번호
	Tel string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}

	return nil
}

// 전화번호의 목록을 출력한다.
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

type Args struct {
	Search string
	List   bool
}

var args = Args{}

func init() {
	flag.BoolVar(&args.List, "list", false, "list entries from phonebook.")
	flag.StringVar(&args.Search, "search", "", "search entries from phonebook.")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s: a phonebook\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	switch {
	// search 커맨드
	case args.Search != "":
		if len(os.Args) != 3 {
			flag.Usage()
			return
		}
		result := search(args.Search)
		if result == nil {
			fmt.Println("Entry not found:", args.Search)
			return
		}
		fmt.Println(*result)
	// list 커맨드
	case args.List:
		list()
	// 커맨드를 찾을 수 없는 경우의 응답
	default:
		fmt.Println("Not a valid option")
	}
}
