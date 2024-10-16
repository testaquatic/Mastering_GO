package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", path.Base(args[0]), "string")
		os.Exit(1)
	}

	fmt.Println(matchInt(args[1]))
}
