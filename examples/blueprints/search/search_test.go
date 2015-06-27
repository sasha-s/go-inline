// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

import "testing"

func TestSearch(t *testing.T) {
	var data = Data{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000}

	var tests = []struct {
		name string
		data Data
		x    Element
		i    int
	}{
		{"empty", nil, 0, 0},
		{"data -20", data, -20, 0},
		{"data -10", data, -10, 0},
		{"data -9", data, -9, 1},
		{"data -6", data, -6, 1},
		{"data -5", data, -5, 1},
		{"data 3", data, 3, 5},
		{"data 11", data, 11, 8},
		{"data 99", data, 99, 9},
		{"data 100", data, 100, 9},
		{"data 101", data, 101, 12},
		{"data 10000", data, 10000, 13},
		{"data 10001", data, 10001, 14},
	}
	for _, e := range tests {
		i := e.data.Search(e.x)
		if i != e.i {
			t.Errorf("%s: expected index %d; got %d", e.name, e.i, i)
		}
	}
}

func TestReverseSearch(t *testing.T) {
	var data = r(Data{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000})

	var tests = []struct {
		name string
		data Data
		x    Element
		i    int
	}{
		{"empty", nil, 0, 0},
		{"data -20", data, -20, 14},
		{"data -10", data, -10, 13},
		{"data -9", data, -9, 13},
		{"data -6", data, -6, 13},
		{"data -5", data, -5, 12},
		{"data 3", data, 3, 8},
		{"data 11", data, 11, 5},
		{"data 99", data, 99, 5},
		{"data 100", data, 100, 2},
		{"data 101", data, 101, 2},
		{"data 10000", data, 10000, 0},
		{"data 10001", data, 10001, 0},
	}
	for _, e := range tests {
		i := e.data.ReverseSearch(e.x)
		if i != e.i {
			t.Errorf("%s: expected index %d; got %d", e.name, e.i, i)
		}
	}
}

// r returns a reversed copy of d.
func r(d Data) Data {
	r := make(Data, len(d))
	for i, x := range d {
		r[len(r)-i-1] = x
	}
	return r
}

// log2 computes the binary logarithm of x, rounded up to the next integer.
// (log2(0) == 0, log2(1) == 0, log2(2) == 1, log2(3) == 2, etc.)
//
func log2(x int) int {
	n := 0
	for p := 1; p < x; p += p {
		// p == 2**n
		n++
	}
	// p/2 < x <= p == 2**n
	return n
}

// Abstract exhaustive test: all sizes up to 100,
// all possible return values.  If there are any small
// corner cases, this test exercises them.
func TestSearchExhaustive(t *testing.T) {
	data := make(Data, 0, 100)
	for size := 0; size <= 100; size++ {
		if size > 0 {
			data = append(data, Element(size-1))
		}
		for targ := 0; targ <= size; targ++ {
			i := data.Search(Element(targ))
			if i != targ {
				t.Errorf("Search(%d, %d) = %d", size, targ, i)
			}
		}
	}
}

// Abstract exhaustive test: all sizes up to 100,
// all possible return values.  If there are any small
// corner cases, this test exercises them.
func TestReverseSearchExhaustive(t *testing.T) {
	data := make(Data, 0, 100)
	for size := 100; size >= 0; size-- {
		if size > 0 {
			data = append(data, Element(size))
		}
		for targ := 100; targ >= size; targ-- {
			i := data.ReverseSearch(Element(targ))
			if i != 100-targ {
				t.Errorf("Search(%d, %d) = %d", size, targ, i)
			}
		}
	}
}
