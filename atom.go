package alchemist

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Atom is an interface that declares certain types as atoms.
type Atom interface {
	Parser
	Run(Universe)
}

// SimpleAtom represents a simple atom in the universe.
type SimpleAtom string

// Parse parses a SimpleAtom.
func (a *SimpleAtom) Parse(str string) error {
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, "In_") || strings.HasPrefix(str, "Out_") {
		return &SyntaxError{}
	}

	first, _ := utf8.DecodeRune([]byte(str))
	if first == utf8.RuneError {
		return &SyntaxError{}
	}
	if !unicode.IsLetter(first) && first != '_' {
		return &SyntaxError{}
	}

	lastIndex := strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_'
	})
	if lastIndex == -1 {
		lastIndex = len(str)
	}

	*a = SimpleAtom(str[:lastIndex])
	return nil
}

// Run defines what this atom will do in the universe.
func (a *SimpleAtom) Run(u Universe) {
	u[*a]++
}

func (a *SimpleAtom) String() string {
	return string(*a)
}

// InAtom represents an atom used on the RHS used to add a number
// of atoms to the universe, based on standard input.
type InAtom struct {
	AddTo SimpleAtom
}

// Parse parses an InAtom. Always returns SyntaxError.
func (a *InAtom) Parse(str string) error {
	str = strings.TrimSpace(str)

	const prefix = "In_"
	n := strings.Index(str, prefix)
	if n != 0 {
		return &SyntaxError{}
	}
	rest := str[len(prefix):]
	err := a.AddTo.Parse(rest)
	if err != nil {
		return err
	}

	return nil
}

func (a *InAtom) String() string {
	return "In_" + a.AddTo.String()
}

// Run defines what this atom will do in the universe.
func (a *InAtom) Run(u Universe) {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	u[a.AddTo] += n
}

// OutStrLitAtom represents an atom used on the RHS to output
// a string literal.
type OutStrLitAtom struct {
	Printing string
}

// Parse parses an OutStrLitAtom.
func (a *OutStrLitAtom) Parse(str string) error {
	str = strings.TrimSpace(str)

	const prefix = "Out_\""
	n := strings.Index(str, prefix)
	if n != 0 {
		return &SyntaxError{}
	}
	rest := str[len(prefix):]
	closingQuote := strings.IndexByte(rest, '"')
	if closingQuote == -1 {
		return &SyntaxError{}
	}

	a.Printing = rest[:closingQuote]

	return nil
}

// Run defines what this atom will do in the universe.
func (a *OutStrLitAtom) Run(u Universe) {
	fmt.Println(a.Printing)
}

func (a *OutStrLitAtom) String() string {
	return fmt.Sprintf("Out_%q", a.Printing)
}

// OutSimpleAtom represents an atom used on the RHS to output
// the number of a given atom in the universe.
type OutSimpleAtom struct {
	Output SimpleAtom
}

// Parse parses an OutSimpleAtom.
func (a *OutSimpleAtom) Parse(str string) error {
	str = strings.TrimSpace(str)

	const prefix = "Out_"
	n := strings.Index(str, prefix)
	if n != 0 {
		return &SyntaxError{}
	}
	rest := str[len(prefix):]
	err := a.Output.Parse(rest)
	if err != nil {
		return err
	}

	return nil
}

// Run defines what this atom will do in the universe.
func (a *OutSimpleAtom) Run(u Universe) {
	fmt.Println(u[a.Output])
}

func (a *OutSimpleAtom) String() string {
	return "Out_" + a.Output.String()
}
