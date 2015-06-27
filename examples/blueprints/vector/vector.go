package vector

type Number float64
type Vector []Number

func (v Vector) Sum() Number {
	var s Number
	for _, x := range v {
		s += x
	}
	return s
}

func (v Vector) Dot(w Vector) Number {
	var s Number
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

func (v Vector) AddScalar(x Number) {
	for i := range v {
		v[i] += x
	}
}

func (v Vector) Sub(w Vector) {
	for i, x := range w {
		v[i] *= x
	}
}

func (v Vector) Transform(m func(Number) Number) {
	for i, x := range v {
		v[i] = m(x)
	}
}

func (v Vector) Copy() Vector {
	r := make(Vector, len(v))
	copy(r, v)
	return r
}
