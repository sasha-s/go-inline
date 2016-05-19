// Package goinline implements inlining for go identifiers.
package goinline

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// Inline replaces each instance of identifier k with v.Ident in ast.File f,
// for k, v := range m.
// For all inlines that were triggeres it also adds imports from v.Imports to f.
// In addition, it removes top level type declarations of the form
// type k ...
// for all k in m.
//
// Every k in m should be a valid identifier.
// Every v.Ident should be a valid expression.
func Inline(fset *token.FileSet, f *ast.File, m map[string]Target) error {
	// Build the inline map.
	im := map[string]reflect.Value{}
	for k, v := range m {
		expr, err := parser.ParseExpr(k)
		if err != nil {
			return fmt.Errorf("failed to parse `%s`: %s", k, err)
		}
		if _, ok := expr.(*ast.Ident); !ok {
			return fmt.Errorf("expected identifier, got %s which is %T", k, expr)
		}
		expr, err = parser.ParseExpr(v.Ident)
		if err != nil {
			return fmt.Errorf("failed to parse `%s`: %s", v.Ident, err)
		}
		s := v.Ident
		if _, ok := expr.(*ast.StarExpr); ok {
			s = fmt.Sprintf("(%s)", s)
		}
		im[k] = reflect.ValueOf(ast.Ident{Name: s})
	}
	// Filter `type XXX ...` declarations out if we are inlining XXX.
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	to := 0
	for _, d := range f.Decls {
		skip := false
		if t, ok := d.(*ast.GenDecl); ok {
			for _, s := range t.Specs {
				ts, ok := s.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if t, ok := m[ts.Name.String()]; ok {
					if !t.NoFiltering {
						skip = true
					}
				}
			}
		}
		if !skip {
			f.Decls[to] = d
			to++
		}
	}
	if to != len(f.Decls) {
		f.Decls = f.Decls[:to]
		// Remove comments for the declarations that were filtered out.
		f.Comments = cmap.Filter(f).Comments()
	}
	// Add imports for the inlines that were triggered.
	for k := range inline(im, f) {
		for _, imp := range m[k].Imports {
			astutil.AddImport(fset, f, imp)
		}
	}
	return nil
}

// Target for inlining.
type Target struct {
	// Ident is a go identifier for the target.
	Ident string
	// Imports are the imporst to be added if the inline is triggered.
	Imports []string
	// NoFiltering prevents removing type Ident when inlining.
	NoFiltering bool
}

// ParseTarget parses a target string.
// Expected format:
// xxx->[import1,import2,...importn::]yyy
// Examples:
// Value->int
// X->go/token::*token.FileSet
func ParseTarget(s string) (string, Target, error) {
	ps := strings.Split(s, "->")
	if len(ps) != 2 {
		return "", Target{}, fmt.Errorf("expected xxx->yyy, got `%s`", s)
	}
	name := ps[0]
	parts := strings.Split(ps[1], "::")
	if len(parts) > 2 {
		return "", Target{}, fmt.Errorf("expected something like a,b,c::v , got `%s`", ps[1])
	}
	if len(parts) == 1 {
		return name, Target{Ident: parts[0]}, nil
	}
	return name, Target{Ident: parts[1], Imports: strings.Split(parts[0], ",")}, nil
}

func inline(im map[string]reflect.Value, node ast.Node) map[string]bool {
	i := inliner{im: im, triggered: map[string]bool{}}
	ast.Walk(i, node)
	return i.triggered
}

type inliner struct {
	im        map[string]reflect.Value
	triggered map[string]bool
}

func (x inliner) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.Ident:
		v, ok := x.im[t.Name]
		if !ok {
			break
		}
		x.triggered[t.Name] = true
		reflect.ValueOf(t).Elem().Set(v)
	}
	return x
}
