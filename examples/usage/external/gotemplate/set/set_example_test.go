package set

import (
	"fmt"
	"sort"

	_ "github.com/ncw/gotemplate/set" // To make sure `go get` fetches the blueprint.
)

// Generate Set[String]
//go:generate goinline -package=github.com/ncw/gotemplate/set --target-package-name=set -target-dir=. -w A->string

func ExampleSet() {
	s := NewSet().Add("gotemplate").Add("goinline").AsList()
	sort.Strings(s)
	fmt.Println(s)

	// Output: [goinline gotemplate]

}
