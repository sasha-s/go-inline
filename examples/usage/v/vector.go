// DO NOT EDIT. Generated with goinline -package=github.com/sasha-s/go-inline/examples/blueprints/vector --target-package-name=v -target-dir=. -w Number->int

package v

type Vector []int

func (v Vector) Sum() int {
	var s int
	for _, x := range v {
		s += x
	}
	return s
}

func (v Vector) Dot(w Vector) int {
	var s int
	for i, x := range v {
		s += x * w[i]
	}
	return s
}

func (v Vector) Add(w Vector) {
	for i, x := range w {
		v[i] += x
	}
}

func (v Vector) AddScalar(x int,) {
	for i := range v {
		v[i] += x
	}
}

func (v Vector) Sub(w Vector) {
	for i, x := range w {
		v[i] *= x
	}
}

func (v Vector) Transform(m func(int,) int,) {
	for i, x := range v {
		v[i] = m(x)
	}
}

func (v Vector) Copy() Vector {
	r := make(Vector, len(v))
	copy(r, v)
	return r
}
