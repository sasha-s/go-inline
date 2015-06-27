package s

import "fmt"

//go:generate goinline -package=github.com/sasha-s/go-inline/examples/blueprints/search --target-package-name=s -target-dir=. -w Element->string

func ExampleSearch() {
	d := Data{"alpha", "beta", "beta", "gamma"}
	fmt.Println(d.Search("alpha"))
	fmt.Println(d.Search("beta"))
	fmt.Println(d.Search("gamma"))
	fmt.Println(d.Search(""))
	fmt.Println(d.Search("z"))

	// Output: 0
	// 1
	// 3
	// 0
	// 4

}

func ExampleReverseSearch() {
	d := Data{"gamma", "beta", "beta", "alpha"}
	fmt.Println(d.ReverseSearch("alpha"))
	fmt.Println(d.ReverseSearch("beta"))
	fmt.Println(d.ReverseSearch("gamma"))
	fmt.Println(d.ReverseSearch(""))
	fmt.Println(d.ReverseSearch("z"))

	// Output: 3
	// 1
	// 0
	// 4
	// 0

}

func ExampleSearchStringSlice() {
	d := Data([]string{"alpha", "beta", "beta", "gamma"})

	fmt.Println(d.Search("alpha"))
	fmt.Println(d.Search("beta"))
	fmt.Println(d.Search("gamma"))
	fmt.Println(d.Search(""))
	fmt.Println(d.Search("z"))

	// Output: 0
	// 1
	// 3
	// 0
	// 4

}
