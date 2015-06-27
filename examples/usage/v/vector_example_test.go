package v

import "fmt"

// Generate Vector[int]
//go:generate goinline -package=github.com/sasha-s/go-inline/examples/blueprints/vector --target-package-name=v -target-dir=. -w Number->int

func ExampleVector() {
	u := make(Vector, 5)
	for i := range u {
		u[i] = i
	}

	u.Transform(func(x int) int {
		return x * x
	})

	fmt.Println(u)
	w := u.Copy()
	w.AddScalar(-4)
	fmt.Println(w)

	fmt.Println(u.Dot(w))
	// Output: [0 1 4 9 16]
	// [-4 -3 0 5 12]
	// 234
}
