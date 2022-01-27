package stringset

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	s := New()
	s.Add("foo")
	s.Add("bar", "baz")

	fmt.Println(s.Contains("foo"))

	s.Remove("bar")
	s.Remove("baz")
	fmt.Println(s)

	// Output:
	// true
	// [foo]
}

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
	s.Add("biz", "baz")
	expected := Set(map[string]struct{}{
		"foo": struct{}{},
		"bar": struct{}{},
		"biz": struct{}{},
		"baz": struct{}{},
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
	assert.ElementsMatch(t, expected, actual)
	// Slice should return an empty slice for an uninitialized set.
	var uninitialized Set
	assert.Exactly(t, []string{}, uninitialized.Slice())
}

func TestSortedSlice(t *testing.T) {
	s := NewFromSlice([]string{"foo", "bar", "foo"})
	expected := []string{"bar", "foo"}
	actual := s.SortedSlice()
	sort.Strings(expected)
	assert.Exactly(t, expected, actual)
	// Slice should return an empty slice for an uninitialized set.
	var uninitialized Set
	assert.Exactly(t, []string{}, uninitialized.Slice())
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		a        Set
		b        Set
		expected Set
	}{
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        NewFromSlice([]string{"c", "d", "e", "f"}),
			expected: NewFromSlice([]string{"a", "b", "c", "d", "e", "f"}),
		},
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        New(),
			expected: NewFromSlice([]string{"a", "b", "c", "d"}),
		},
		{
			a:        New(),
			b:        NewFromSlice([]string{"a", "b", "c", "d"}),
			expected: NewFromSlice([]string{"a", "b", "c", "d"}),
		},
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        nil,
			expected: NewFromSlice([]string{"a", "b", "c", "d"}),
		},
		{
			a:        nil,
			b:        NewFromSlice([]string{"a", "b", "c", "d"}),
			expected: NewFromSlice([]string{"a", "b", "c", "d"}),
		},
		{
			a:        nil,
			b:        nil,
			expected: New(),
		},
		{
			a:        New(),
			b:        New(),
			expected: New(),
		},
	}
	for i, tc := range testCases {
		actual := Union(tc.a, tc.b)
		assert.Exactly(t, tc.expected, actual, "test case: %d", i)
	}
}

func TestIntersect(t *testing.T) {
	testCases := []struct {
		a        Set
		b        Set
		expected Set
	}{
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        NewFromSlice([]string{"c", "d", "e", "f"}),
			expected: NewFromSlice([]string{"c", "d"}),
		},
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        New(),
			expected: New(),
		},
		{
			a:        New(),
			b:        NewFromSlice([]string{"a", "b", "c", "d"}),
			expected: New(),
		},
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        nil,
			expected: New(),
		},
		{
			a:        nil,
			b:        NewFromSlice([]string{"a", "b", "c", "d"}),
			expected: New(),
		},
		{
			a:        nil,
			b:        nil,
			expected: New(),
		},
		{
			a:        New(),
			b:        New(),
			expected: New(),
		},
	}
	for i, tc := range testCases {
		actual := Intersect(tc.a, tc.b)
		assert.Exactly(t, tc.expected, actual, "test case: %d", i)
	}
}

func TestDiff(t *testing.T) {
	testCases := []struct {
		a        Set
		b        Set
		expected Set
	}{
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        NewFromSlice([]string{"c", "d", "e", "f"}),
			expected: NewFromSlice([]string{"a", "b"}),
		},
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        New(),
			expected: NewFromSlice([]string{"a", "b", "c", "d"}),
		},
		{
			a:        New(),
			b:        NewFromSlice([]string{"a", "b", "c", "d"}),
			expected: New(),
		},
		{
			a:        NewFromSlice([]string{"a", "b", "c", "d"}),
			b:        nil,
			expected: NewFromSlice([]string{"a", "b", "c", "d"}),
		},
		{
			a:        nil,
			b:        NewFromSlice([]string{"a", "b", "c", "d"}),
			expected: New(),
		},
		{
			a:        nil,
			b:        nil,
			expected: New(),
		},
		{
			a:        New(),
			b:        New(),
			expected: New(),
		},
	}
	for i, tc := range testCases {
		actual := Diff(tc.a, tc.b)
		assert.Exactly(t, tc.expected, actual, "test case: %d", i)
	}
}
