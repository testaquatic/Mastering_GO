package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s datetime\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	dateTime, err := time.Parse("02 January 2006 15:04 MST", flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	fmt.Println("Current Location:", dateTime.Local())
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Println("New York Time:", dateTime.In(loc))
	fmt.Println("London time:", dateTime.UTC())
	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Println("Tokyo Time:", dateTime.In(loc))
}
