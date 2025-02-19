package quickT

import (
	"testing"
	"testing/quick"
)

// 이보다 큰 수를 사용해야 타임아웃 될 수도 있다.
const N = 1_000_000

func TestWithItself(t *testing.T) {
	condition := func(a, b Point2D) bool {
		return Add(a, b) == Add(b, a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestThree(t *testing.T) {
	condition := func(a, b, c Point2D) bool {
		return Add(Add(a, b), c) == Add(a, b)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
