// Adapted from https://github.com/golang/go/tree/master/src/sort/search.go

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://github.com/golang/go/blob/master/LICENSE

// This file implements binary search for native types.

package search

// Element is an element type.
type Element int // Can be anything comparable and ordered. See http://golang.org/ref/spec#Comparison_operators.

// Data is a slice of elements.
type Data []Element

// Len returns a lenth of a slice.
func (data Data) Len() int {
	return len(data)
}

// Search uses binary search to find and return the smallest index i
// in [0, n) at which with data[i] >= x is true, assuming that data is sorted in ascending order.
// If there is no such index, Search returns n.
func (data Data) Search(x Element) int {
	// f(i) == data[i] >= x
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(data)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if data[h] < x {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

// ReverseSearch uses binary search to find and return the smallest index i
// in [0, n) at which with data[i] <= x is true, assuming that data is sorted in descending order.
// If there is no such index, Search returns n.
func (data Data) ReverseSearch(x Element) int {
	// f(i) == data[i] <= x
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(data)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if data[h] > x {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
