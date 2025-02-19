package testRE

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test_matchInt(t *testing.T) {
	if matchInt("") {
		t.Error(`matchInt("") != false`)
	}

	if !matchInt("00") {
		t.Error(`matchInt("00") != true`)
	}

	if !matchInt("-00") {
		t.Error(`matchInt("-00") != true`)
	}

	if !matchInt("+00") {
		t.Error(`matchInt("+00") != true`)
	}
}

func Test_with_random(t *testing.T) {
	n := strconv.Itoa(random(-100_000, 19_999))

	if !matchInt(n) {
		t.Errorf(`matchInt(%s) != false`, n)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
