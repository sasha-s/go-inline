// DO NOT EDIT. Generated with goinline -package=github.com/sasha-s/go-inline/examples/blueprints/concurrentmap --target-package-name=cm_string_int -target-dir=cm_string_int -w Value->int Key->string

// Package concurrentmap implements typesafe map safe for concurrent use.
package cm_string_int

import "sync"

// CM is a concurrent map from Key to Value.
// A `map[Key]Value` with a lock.
type CM struct {
	mu	sync.RWMutex
	m	map[string]int
}

// New creates a new CM, large enough to accomodate hint inserts without resizing.
func New(hint int) *CM {
	return &CM{m: make(map[string]int, hint)}
}

// Get returns (value, true) if k is in a map,
// (zero value, false) otherwise.
// Similar to `v, ok := m[k]`.
func (m *CM) Get(k string,) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

// Set sets the value of m[k] to v.
// Similar to `m[k] = v`.
// Returns (v, true) if the value was inserted, (old value, false) otherwise.
func (m *CM) Set(k string, v int,) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[k] = v
}

// Insert a (k, v) pair into a map if it is not already there.
// Returns (v, true) if the value was inserted, (old value, false) otherwise.
func (m *CM) Insert(k string, v int,) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	old, ok := m.m[k]
	if ok {
		return old, false
	}
	m.m[k] = v
	return v, true
}

// InsertF inserts a (k, f()) pair into a map if it is not already there.
// Returns (new value, true) if the value was inserted, (old value, false) otherwise.
// InsertF does not call f() if is not needed.
// Useful when constructing a new Value is expensive.
func (m *CM) InsertF(k string, f func() int,) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	old, ok := m.m[k]
	if ok {
		return old, false
	}
	v := f()
	m.m[k] = v
	return v, true
}

// Remove value for a key k if present, a no-op otherwise.
// Similar to delete(m, key).
func (m *CM) Remove(k string,) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, k)
}

// Len returns number of elements in a map.
// Similar to `len(m)`.
func (m *CM) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.m)
}

// Keys returns a slice containing all the keys in the map.
func (m *CM) Keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]string, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}
