package alchemist

import (
	"strconv"
	"strings"
	"unicode"
)

// Rule represents a rule in the program.
type Rule struct {
	Left  LHSRule
	Right RHSRule
}

// LHSRule represents the left-hand side of a rule.
type LHSRule map[SimpleAtom]int

// Parse allows an LHSRule to be parsed from a string.
//
// Format (whitespace optional): `\d* (SimpleAtom) (\+ \d* SimpleAtom)*`
func (r *LHSRule) Parse(str string) error {
	str = strings.TrimSpace(str)

	rule := make(LHSRule)
	*r = rule

	rest := str
	var index int
	for {
		rest = rest[index:]
		firstNonDigit := strings.IndexFunc(rest, func(r rune) bool {
			return !unicode.IsDigit(r)
		})

		var num int
		if firstNonDigit > 0 {
			digits := rest[:firstNonDigit]
			num, _ = strconv.Atoi(digits)
		} else {
			num = 1
		}

		var atom SimpleAtom
		err := atom.Parse(rest[firstNonDigit:])
		if err != nil {
			return &SyntaxError{
				reason: "unable to parse an atom in " + rest,
			}
		}

		if num == 0 {
			rule[atom] = 0
		}
		for i := 0; i < num; i++ {
			rule[atom]++
		}

		next := strings.IndexRune(rest, '+')
		if next < 0 {
			break
		} else if next == len(rest)-1 {
			return &SyntaxError{
				reason: "trailing + sign in " + str,
			}
		} else {
			index = next + 1
		}
	}

	return nil
}

// RHSRule represents the right-hand side of a rule.
type RHSRule []Atom

// Parse allows an LHSRule to be parsed from a string.
//
// Format (whitespace optional): `\d* (Atom) (\+ \d* Atom)*`
func (r *RHSRule) Parse(str string) error {
	str = strings.TrimSpace(str)

	rest := str
	var index int
	for {
		rest = rest[index:]
		firstNonDigit := strings.IndexFunc(rest, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		var num int
		if firstNonDigit > 0 {
			digits := rest[:firstNonDigit]
			num, _ = strconv.Atoi(digits)
		} else {
			num = 1
		}

		var simple SimpleAtom
		var in InAtom
		var outSimple OutSimpleAtom
		var outStrLit OutStrLitAtom

		var correct Atom

		atoms := []Atom{&simple, &in, &outSimple, &outStrLit}
		for _, a := range atoms {
			err := a.Parse(rest)
			if err == nil {
				correct = a
				break
			}
		}

		for i := 0; i < num; i++ {
			*r = append(*r, correct)
		}

		next := strings.IndexRune(rest, '+')
		if next < 0 {
			break
		} else if next == len(rest)-1 {
			return &SyntaxError{
				reason: "trailing + sign in " + str,
			}
		} else {
			index = next + 1
		}
	}

	return nil

}
