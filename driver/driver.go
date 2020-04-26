package driver

import (
	"io"
)

// Coverage //
type Coverage struct {
	Package string
	Pecent  float32
	Order   int // ordinal position within a coeverage run
}

// Set is a set of coverages, keyed by Coverage.Package
type Set map[string]Coverage

// Parser is an interface that can parse the code coverage output of a given language
type Parser interface {
	Parse(io.Reader) (Set, error)
}
