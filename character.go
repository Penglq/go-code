package main

import (
	"log"
	"unicode"
)

type CharSpec interface {
	Satisfy(c rune) bool
}

type HasDigit struct{}

func (I *HasDigit) Satisfy(c rune) bool {
	return unicode.IsDigit(c)
}

type HasUpper struct{}

func (I HasUpper) Satisfy(c rune) bool {
	return unicode.IsUpper(c)
}

type Equals struct {
	c rune
}

func (e Equals) Satisfy(c rune) bool {
	return e.c == c
}

func Exists(word string, spec CharSpec) bool {
	return Expect(word, spec, true)
}

type IsNot struct {
	spec CharSpec
}

func (i IsNot) Satisfy(c rune) bool {
	return !i.spec.Satisfy(c)
}

func ForAll(word string, spec CharSpec) bool {
	return Expect(word, spec, false)
}

func Expect(word string, spec CharSpec, ok bool) bool {
	for _, c := range word {
		if spec.Satisfy(c) == ok {
			return ok
		}
	}
	return !ok
}

type AnyOf struct {
	specs []CharSpec
}

func (a AnyOf) Satisfy(c rune) bool {
	for _, spec := range a.specs {
		if spec.Satisfy(c) {
			return true
		}
	}
	return false
}
func main() {
	log.Print(ForAll("ab12", IsNot{Equals{'1'}}))
	log.Print(Exists("ab12", AnyOf{[]CharSpec{Equals{'g'}, Equals{'a'}}}))
}
