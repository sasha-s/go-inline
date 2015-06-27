package example_sort

import (
	"fmt"
	"math"

	"github.com/sasha-s/go-inline/examples/usage/sort/float64_sort"
	"github.com/sasha-s/go-inline/examples/usage/sort/string_sort"
)

//go:generate goinline -w -package github.com/sasha-s/go-inline/examples/blueprints/timsort --target-package-name=string_sort -target-dir=string_sort Value->string

//go:generate goinline -w -package github.com/sasha-s/go-inline/examples/blueprints/timsort --target-package-name=float64_sort -target-dir=float64_sort Value->float64

func ExampleSortStrings() {
	data := []string{"alpha", "1", "beta", "2"}
	string_sort.Sort(data)
	fmt.Println(data)

	// Output:  [1 2 alpha beta]

}

func ExampleReverseSortStrings() {
	data := []string{"alpha", "1", "beta", "2"}
	string_sort.ReverseSort(data)
	fmt.Println(data)

	// Output:  [beta alpha 2 1]

}

func ExampleSortFloats() {
	data := []float64{1, 0, 3.14159265359, 2.71828182846, 1.41421356237, 1.61803398875, math.Inf(1), math.Inf(-1)}
	float64_sort.Sort(data)
	fmt.Println(data)

	// Output:  [-Inf 0 1 1.41421356237 1.61803398875 2.71828182846 3.14159265359 +Inf]

}

func ExampleReverseSortFloats() {
	data := []float64{1, 0, 3.14159265359, 2.71828182846, 1.41421356237, 1.61803398875, math.Inf(1), math.Inf(-1)}
	float64_sort.ReverseSort(data)
	fmt.Println(data)

	// Output:  [+Inf 3.14159265359 2.71828182846 1.61803398875 1.41421356237 1 0 -Inf]

}
