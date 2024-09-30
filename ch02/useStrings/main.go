package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAlis"))
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Mihalis", "MI"))

	fmt.Printf("Index: %v\n", strings.Index("Mihalis", "ha"))
	fmt.Printf("Index: %v\n", strings.Index("Mihalis", "HA"))

	fmt.Printf("Prefix: %v\n", strings.HasPrefix("Mihalis", "Mi"))
	fmt.Printf("Prefix: %v\n", strings.HasPrefix("Mihalis", "mi"))
	fmt.Printf("Suffix: %v\n", strings.HasSuffix("Mihalis", "is"))
	fmt.Printf("Suffix: %v\n", strings.HasSuffix("Mihalis", "IS"))

	t := strings.Fields("This is a string!")
	fmt.Printf("Fields: %v\n", len(t))
	t = strings.Fields("ThisIs a\tstring")
	fmt.Printf("Fields: %v\n", len(t))

	fmt.Printf("%s\n", strings.Split("abcd efg", ""))
	fmt.Printf("%s\n", strings.Replace("abcd efg", "", "_", -1))
	fmt.Printf("%s\n", strings.Replace("abcd efg", "", "_", 4))
	fmt.Printf("%s\n", strings.Replace("abcd efg", "", "_", 2))

	fmt.Printf("SplitAter: %s\n", strings.SplitAfter("123++432++", "++"))

	trimFuction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("TrimFunc: %s\n", strings.TrimFunc("123 abc ABC \t .", trimFuction))
}
