package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
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

func search(key string) int64 {
	count := int64(0)
	for i, v := range data {
		if v.Tel == key {
			count += 1
			fmt.Println(&data[i])
		}
	}
	return count
}

// 전화번호의 목록을 출력한다.
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

const (
	MIN = 0
	MAX = 94
)

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		data = append(data, Entry{name, surname, n})
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

	populate(100)

	fmt.Println("Data has", len(data), "entries.")

	switch {
	// search 커맨드
	case args.Search != "":
		if len(os.Args) != 3 {
			flag.Usage()
			return
		}
		result := search(args.Search)
		if result == 0 {
			fmt.Println("Entry not found:", args.Search)
			return
		}
	// list 커맨드
	case args.List:
		list()
	// 커맨드를 찾을 수 없는 경우의 응답
	default:
		fmt.Println("Not a valid option")
	}
}
