package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

func zero5() Entry {
	return Entry{}
}

func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	s1 := zero5()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := initS("Nihalis", "Tsoukalos", 2020)
	p2 := initPtoS("Nihalis", "Tsoukalos", 2020)
	fmt.Println("s2:", s2, "p2:", *p2)

	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)

	pS := new(Entry)
	fmt.Println("pS:", pS)
}
