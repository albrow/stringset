# stringset

[![Version](https://img.shields.io/badge/version-2.1.0-5272B4.svg)](https://github.com/albrow/stringset/releases)
[![Circle CI](https://img.shields.io/circleci/project/albrow/stringset/master.svg)](https://circleci.com/gh/albrow/stringset/tree/master)
[![GoDoc](https://godoc.org/github.com/albrow/stringset?status.svg)](https://godoc.org/github.com/albrow/stringset)

A simple and space-effecient Go implementation of a set of strings.

## Basic Usage

```go
s := stringset.New()
s.Add("foo")
s.Add("bar", "baz")
fmt.Println(s)
// Output:
// [bar foo baz]

fmt.Println(s.Contains("foo"))
// Output:
// true

s.Remove("bar")
fmt.Println(s)
// Output:
// [foo]
```
