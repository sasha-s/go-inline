package timsort

import (
	"fmt"
	"math/rand"
	"testing"
)

func makeTestArrayI(size int) []Value {
	a := make([]Value, size)

	for i := 0; i < size; i++ {
		a[i] = Value(i & 0xeeeeee)
	}

	return a
}

func IsSortedI(a []Value) bool {
	len := len(a)

	if len < 2 {
		return true
	}

	prev := a[0]
	for i := 1; i < len; i++ {
		if a[i] < prev {
			fmt.Println("false")
			return false
		}
	}

	return true
}

func TestSmokeI(t *testing.T) {
	a := []Value{3, 1, 2}

	Sort(a)

	if !IsSortedI(a) {
		t.Error("not sorted")
	}
}

func Test1KI(t *testing.T) {
	a := makeTestArrayI(1024)

	Sort(a)
}

func Test1MI(t *testing.T) {
	a := makeTestArrayI(1024 * 1024)

	Sort(a)
	if !IsSortedI(a) {
		t.Error("not sorted")
	}
}

func makeRandomArrayI(size int) []Value {
	a := make([]Value, size)

	for i := 0; i < size; i++ {
		a[i] = Value(rand.Intn(100))
	}

	return a
}

func TestRandom1MI(t *testing.T) {
	size := 1024 * 1024

	a := makeRandomArrayI(size)
	b := make([]Value, size)
	copy(b, a)

	Sort(a)
	if !IsSortedI(a) {
		t.Error("not sorted")
	}
}
