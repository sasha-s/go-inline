# go-inline
Generic Data Structures/Algorithms in golang via code generation (glorified copy&paste).

One might one to use goinline tool to generate specific implementation of generic data structures, such as map, vector, matrix, set.
Or to  generate specific implementation of generic algorithms, such as sort, binary search.

Goals:

* Write code once
* Simple, readable and tested tested go code (e.g. no templates/language extensions)
* Type safety and speed all the way (e.g. no reflection)
* No magic, everything is explicit

##### How-to
Install goinline.
```
go get github.com/sasha-s/go-inline/goinline/cmd/goinline
```

* [Create a blueprint](#creating-a-blueprint). Some [examples](https://github.com/sasha-s/go-inline/blob/master/examples/blueprints)
* [Generate derived code](#generating-derived-code). Some [examples](https://github.com/sasha-s/go-inline/blob/master/examples/usage). Look for go:generate directives.
* Use the generated code.  Some [examples](https://github.com/sasha-s/go-inline/blob/master/examples/usage)
* Check the generated code in.


##### Creating a blueprint.

Blueprint is a package that implements a certain generic data type/algorithm.

Let us consider Vector[Number].

```go
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
...
```
Number is a type parameter: we want to be able to use derived versions Vector[float32] or Vector[int].

Note, that Number is a standalone type. This makes it easy to generate the derived versions.

Since we are using `+` and `*` the Vector blueprint can be only used with ints and floats.

Have a look at more complete [example](https://github.com/sasha-s/go-inline/blob/master/examples/blueprints/vector/vector.go)

##### Generating derived code.

In case of Vector[Number]:

```
goinline -package=github.com/sasha-s/go-inline/examples/blueprints/vector --target-package-name=v -target-dir=. -w "Number->int"
```

This will use a blueprint from package `github.com/sasha-s/go-inline/examples/blueprints/vector`, looking for the it in `$GOPATH`, create a generated version in current (`.`) folder, will rename the package to `v` and will replace `Number` with `int` in the generated code.

Equivalent [go:generate](http://blog.golang.org/generate) directive:
//go:generate goinline -package=github.com/sasha-s/go-inline/examples/blueprints/vector --target-package-name=v -target-dir=. -w Number->int

See [vector example](https://github.com/sasha-s/go-inline/blob/master/examples/usage/v/vector_example_test.go).

Note, goinline does not check if the blueprint code compiles. Garbage in, garbage out.

Writing tests so they work after inlining types is tricky, so goinline does not generate tests by default.

##### goinline tool

<pre>
goinline -h
Usage of goinline:
  -package="": package to use as a blueprint. Something like `github.com/sasha-s/go-inline/examples/blueprints/concurrentmap`
  -target-dir="": where to put the generated code. Will modify the blueprint (according to package) if empty
  -target-package-name="": package name for the generated code. Ignored if empty
  -tests=false: process tests as well
  -w=false: write result to a file instead of stdout
</pre>

One could get similar results with some scripting around [gofmt -r](https://golang.org/cmd/gofmt/), given that the blueprints are well-structured.

##### FAQ
* Why is there no generics in go?
 - [Official answer](https://golang.org/doc/faq#generics)
 - [Detailed analisys](https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/)
* How is it different from
 - copy&paste: c&p is painful to maintain, though if the blueprint is well-structured, the copy&paste is easy.
 - [gen](http://clipperhouse.github.io/gen/): gen works with [text/template](https://github.com/clipperhouse/linkedlist/blob/master/templates.go), goinline starts with [working, testable go code](https://github.com/sasha-s/go-inline/blob/master/examples/blueprints/search/search.go), so creating the blueprints is easier and cleaner.
 - [goast](https://github.com/go-goast/goast): goast tries to be smart and infer things, goinline is very explicit. Also, goinline code is shorter and simpler.
 - [gotemplate](https://github.com/ncw/gotemplate): goinline is more explicit, works with multiple files per blueprint. More [details](https://groups.google.com/d/msg/golang-nuts/8wwQcaGRVD4/EfGOpx4A3igJ).

