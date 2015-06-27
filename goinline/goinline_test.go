package goinline

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"testing"
)

var file = []byte(
	`// Comment 1
package example
// Comment 2
type X struct {
// Comment 3
}
// Comment 4
type Y X
// Comment 5
type Z []X
// Comment 6
func f (x X) Z {
	var _ Z
	return Z(nil)
}
`)

func TestInline(t *testing.T) {
	var fset = token.NewFileSet()
	f, err := parser.ParseFile(fset, "data.go", file, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	err = Inline(fset, f, map[string]Target{"X": Target{
		Ident:   "*int",
		Imports: []string{"a.b", "c.d"},
	}})
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	expected := `// Comment 1
package example

import (
	"a.b"
	"c.d"
)

// Comment 3

// Comment 4
type Y (*int)

// Comment 5
type Z [](*int)

// Comment 6
func f(x (*int),) Z {
	var _ Z
	return Z(nil)
}`
	got := strings.TrimSpace(buf.String())
	expected = strings.TrimSpace(expected)
	if expected != got {
		for i := 0; i < len(expected) && i < len(got); i++ {
			if expected[i] != got[i] {
				t.Errorf("expected:...\n%s\ngot:...\n%s\n", expected[i:], got[i:])
				break
			}
		}
		if len(expected) != len(got) {
			t.Errorf("expected: %d, got %d", len(expected), len(got))
		}
	}
}

func TestInline2(t *testing.T) {
	var fset = token.NewFileSet()
	f, err := parser.ParseFile(fset, "data.go", file, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	err = Inline(fset, f, map[string]Target{"Z": Target{
		Ident: "[]int",
	}})
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	expected := `// Comment 1
package example

// Comment 2
type X struct {
	// Comment 3
}

// Comment 4
type Y X

// Comment 6
func f(x X) []int {
	var _ []int
	return []int(nil)
}`
	got := strings.TrimSpace(buf.String())
	expected = strings.TrimSpace(expected)
	if expected != got {
		for i := 0; i < len(expected) && i < len(got); i++ {
			if expected[i] != got[i] {
				t.Errorf("expected:...\n%s\ngot:...\n%s\n", expected[i:], got[i:])
				break
			}
		}
		if len(expected) != len(got) {
			t.Errorf("expected: %d, got %d", len(expected), len(got))
		}
	}
}

func TestInlineBadMap(t *testing.T) {
	err := s(Inline(nil, nil, map[string]Target{"-Z": Target{
		Ident: "[]int",
	}}))
	expected := `expected identifier, got -Z which is *ast.UnaryExpr`
	if err != expected {
		t.Errorf("expected [%v], got [%v]", expected, err)
	}
}

func TestInlineBadMap2(t *testing.T) {
	err := s(Inline(nil, nil, map[string]Target{"a.x": Target{
		Ident: "[]int",
	}}))
	expected := `expected identifier, got a.x which is *ast.SelectorExpr`
	if err != expected {
		t.Errorf("expected [%v], got [%v]", expected, err)
	}
}

func TestInlineBadMap3(t *testing.T) {
	err := s(Inline(nil, nil, map[string]Target{"0x": Target{
		Ident: "[]int",
	}}))
	expected := "failed to parse `0x`: 1:1: illegal hexadecimal number"
	if err != expected {
		t.Errorf("expected [%v], got [%v]", expected, err)
	}
}

func TestInlineBadMap4(t *testing.T) {
	err := s(Inline(nil, nil, map[string]Target{"x": Target{
		Ident: "[-]int",
	}}))
	expected := "failed to parse `[-]int`: 1:3: expected operand, found ']'"
	if err != expected {
		t.Errorf("expected [%v], got [%v]", expected, err)
	}
}

func TestParseTarget(t *testing.T) {
	for _, tc := range []struct {
		in   string
		err  string
		name string
		t    Target
	}{
		{
			in:  "",
			err: "expected xxx->yyy, got ``",
		},
		{
			in: "->",
		},
		{
			in:   "x->y",
			name: "x",
			t: Target{
				Ident: "y",
			},
		},
		{
			in:  "x->y->z",
			err: "expected xxx->yyy, got `x->y->z`",
		},
		{
			in:   "x->p::y",
			name: "x",
			t: Target{
				Ident:   "y",
				Imports: []string{"p"},
			},
		},
		{
			in:  "x->p::y::z",
			err: "expected something like a,b,c::v , got `p::y::z`",
		},
		{
			in:   "alpha->golang.org/x/tools/go/ast/astutil,go/token::*token.FileSet",
			name: "alpha",
			t: Target{
				Ident:   "*token.FileSet",
				Imports: []string{"golang.org/x/tools/go/ast/astutil", "go/token"},
			},
		},
	} {
		name, trg, err := ParseTarget(tc.in)
		if tc.name != name {
			t.Errorf("expected [%v], got [%v]", tc.name, name)
		}
		if tc.err != s(err) {
			t.Errorf("expected [%v], got [%v]", tc.err, s(err))
		}
		if str(tc.t) != str(trg) {
			t.Errorf("expected [%v], got [%v]", tc.t, trg)
		}
	}
}

func str(t Target) string {
	return fmt.Sprintf("%#v", t)
}

func s(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
