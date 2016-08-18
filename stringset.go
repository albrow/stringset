// Package stringset is a simple and space-effecient implementation of a set of
// strings.
package stringset

import "fmt"

// Set is an unsorted set of unique strings.
type Set map[string]struct{}

// New returns an initialized Set.
func New() Set {
	return Set{}
}

// NewFromSlice returns a new set constructed from the given slice. Any
// duplicate values will be removed.
func NewFromSlice(slice []string) Set {
	s := New()
	for _, v := range slice {
		s.Add(v)
	}
	return s
}

// Add adds v to the set. If v is alredy in the set, this has no effect.
func (s Set) Add(v string) {
	s[v] = struct{}{}
}

// Remove removes v from the set. If v is not in the set, this has no effect.
func (s Set) Remove(v string) {
	delete(s, v)
}

// Contains returns true if the set contains v and false otherwise.
func (s Set) Contains(v string) bool {
	_, ok := s[v]
	return ok
}

// Slice returns the values in the set as a slice of strings. It returns an
// empty slice if the set contains no values. The values returned will be in
// random order.
func (s Set) Slice() []string {
	slice := make([]string, len(s))
	i := 0
	for v := range s {
		slice[i] = v
		i++
	}
	return slice
}

// String implements the Stringer interface.
func (s Set) String() string {
	return fmt.Sprint(s.Slice())
}
