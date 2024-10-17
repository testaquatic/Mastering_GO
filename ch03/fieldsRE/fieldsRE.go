package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^\d+$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}
	if !matchNameSur(fields[0]) {
		return false
	}
	if !matchNameSur(fields[1]) {
		return false
	}
	return matchTel(fields[2])
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("사용법:", path.Base(os.Args[0]), "문자열")
		os.Exit(1)
	}

	s := flag.Arg(0)

	fmt.Println(matchRecord(s))
}
