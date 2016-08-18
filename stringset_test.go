package stringset

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromSlice(t *testing.T) {
	s := NewFromSlice([]string{"foo", "bar", "foo"})
	expected := Set(map[string]struct{}{
		"foo": struct{}{},
		"bar": struct{}{},
	})
	assert.Exactly(t, expected, s)
}

func TestAdd(t *testing.T) {
	s := New()
	s.Add("foo")
	s.Add("foo")
	s.Add("bar")
	expected := Set(map[string]struct{}{
		"foo": struct{}{},
		"bar": struct{}{},
	})
	assert.Exactly(t, expected, s)
	// Check that calling Add on an uninitialized set panics.
	var uninitialied Set
	assert.Panics(t, func() { uninitialied.Add("foo") })
}

func TestRemove(t *testing.T) {
	s := NewFromSlice([]string{"foo", "bar", "biz"})
	s.Remove("biz")
	expected := Set(map[string]struct{}{
		"foo": struct{}{},
		"bar": struct{}{},
	})
	assert.Exactly(t, expected, s)
}

func TestContains(t *testing.T) {
	s := NewFromSlice([]string{"foo", "bar"})
	assert.True(t, s.Contains("foo"))
	assert.True(t, s.Contains("bar"))
	assert.False(t, s.Contains("biz"))
	// Contains should return false for an uninitialized set.
	var uninitialized Set
	assert.False(t, uninitialized.Contains("foo"))
	assert.False(t, uninitialized.Contains("bar"))
}

func TestSlice(t *testing.T) {
	s := NewFromSlice([]string{"foo", "bar", "foo"})
	expected := []string{"foo", "bar"}
	actual := s.Slice()
	sort.Strings(expected)
	sort.Strings(actual)
	assert.Exactly(t, expected, actual)
	// Slice should return an empty slice for an uninitialized set.
	var uninitialized Set
	assert.Exactly(t, []string{}, uninitialized.Slice())
}

func Example() {
	s := New()
	s.Add("foo")
	s.Add("bar")

	fmt.Println(s.Contains("foo"))

	s.Remove("bar")
	fmt.Println(s)

	// Output:
	// true
	// [foo]
}
