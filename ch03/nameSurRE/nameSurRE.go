package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage:", path.Base(os.Args[0]), "TEXT")
		os.Exit(1)
	}
	s := flag.Arg(0)
	fmt.Println(matchNameSur(s))
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}
