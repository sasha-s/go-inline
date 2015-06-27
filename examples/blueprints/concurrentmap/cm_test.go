package concurrentmap

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	tcs := []struct {
		m  map[Key]Value
		k  Key
		v  Value
		ex Value
		ok bool
	}{
		{m: map[Key]Value{}, k: 1, v: "a", ex: "a", ok: true},
		{m: map[Key]Value{2: "z"}, k: 1, v: "a", ex: "a", ok: true},
		{m: map[Key]Value{1: "z"}, k: 1, v: "a", ex: "z", ok: false},
	}
	for _, tc := range tcs {
		m := &CM{m: tc.m}
		v, ok := m.Insert(tc.k, tc.v)
		if v != tc.ex {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", tc.ex, v, tc)
		}
		if ok != tc.ok {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", tc.ok, ok, tc)
		}
	}
}

func TestInsertF(t *testing.T) {
	tcs := []struct {
		m  map[Key]Value
		k  Key
		v  Value
		ex Value
		ok bool
	}{
		{m: map[Key]Value{}, k: 1, v: "a", ex: "a", ok: true},
		{m: map[Key]Value{2: "z"}, k: 1, v: "a", ex: "a", ok: true},
		{m: map[Key]Value{1: "z"}, k: 1, v: "a", ex: "z", ok: false},
	}
	for _, tc := range tcs {
		m := &CM{m: tc.m}
		v, ok := m.InsertF(tc.k, func() Value {
			return tc.v
		})
		if v != tc.ex {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", tc.ex, v, tc)
		}
		if ok != tc.ok {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", tc.ok, ok, tc)
		}
	}
}

func TestGet(t *testing.T) {
	tcs := []struct {
		m  map[Key]Value
		k  Key
		ex Value
		ok bool
	}{
		{m: map[Key]Value{}, k: 1, ex: "", ok: false},
		{m: map[Key]Value{2: "z"}, k: 1, ex: "", ok: false},
		{m: map[Key]Value{1: "z"}, k: 1, ex: "z", ok: true},
	}
	for _, tc := range tcs {
		m := &CM{m: tc.m}
		v, ok := m.Get(tc.k)
		if v != tc.ex {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", tc.ex, v, tc)
		}
		if ok != tc.ok {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", tc.ok, ok, tc)
		}
	}
}

func TestRemove(t *testing.T) {
	tcs := []struct {
		m  map[Key]Value
		k  Key
		ex map[Key]Value
	}{
		{m: nil, k: 1, ex: nil},
		{m: map[Key]Value{}, k: 1, ex: nil},
		{m: map[Key]Value{1: "z"}, k: 2, ex: map[Key]Value{1: "z"}},
		{m: map[Key]Value{1: "z"}, k: 1, ex: map[Key]Value{}},
		{m: map[Key]Value{1: "z", 2: "a", 3: "d"}, k: 1, ex: map[Key]Value{2: "a", 3: "d"}},
	}
	for _, tc := range tcs {
		m := &CM{m: tc.m}
		m.Remove(tc.k)
		keys := m.Keys()
		if len(keys) != len(tc.ex) {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", len(tc.ex), len(keys), tc)
		}
		if m.Len() != len(tc.ex) {
			t.Fatalf("Expected `%v` got `%v` | %#v\n", len(tc.ex), m.Len(), tc)
		}
		for _, k := range keys {
			v, ok := m.Get(k)
			ex, exOK := tc.ex[k]
			if v != ex {
				t.Fatalf("Expected `%v` got `%v` | %#v\n", ex, v, tc)
			}

			if ok != exOK {
				t.Fatalf("Expected `%v` got `%v` | %#v\n", exOK, ok, tc)
			}
		}
	}
}

func TestRace(t *testing.T) {
	m := New(0)
	for g := 0; g < 1000; g++ {
		go func() {
			for k := 0; k < 50; k++ {
				m.Insert(Key(k), Value(fmt.Sprint(k)))
				m.Get(Key(k))
				m.Keys()
			}
		}()
	}
}
