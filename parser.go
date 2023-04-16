package alchemist

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode"
)

var errNowAtInput = errors.New("now at input")

// Parser is an interface that describes objects which can be parsed from strings.
type Parser interface {
	Parse(string) error
}

// custom bufio.SplitFunc that splits on newlines, arrows (->), and bangs (!).
// newlines and arrows are suppressed from output, but bangs are not.
func splitter(data []byte, atEOF bool) (advance int, token []byte, err error) {

	oldLen := len(data)
	data = bytes.TrimLeftFunc(data, unicode.IsSpace)
	trim := oldLen - len(data)

	var lastByte byte
	var stringMode bool
	for i, b := range data {
		if lastByte != '\\' && b == '"' {
			stringMode = !stringMode
		}
		if stringMode {
			continue
		}
		if b == '!' {
			trimmed := bytes.TrimSpace(data[:i+1])
			return i + 1 + trim, trimmed, nil
		}
		if lastByte == '-' && b == '>' {
			trimmed := bytes.TrimSpace(data[:i-1])
			return i + 1 + trim, trimmed, nil
		}
		if b == '\n' {
			trimmed := bytes.TrimSpace(data[:i])
			return i + 1 + trim, trimmed, nil
		}
		lastByte = b
	}

	if atEOF {
		trimmed := bytes.TrimSpace(data)
		return len(data) + trim, trimmed, nil
	}

	return 0, nil, nil
}

// Parse reads the bytes from r and parses them
// into a returned array of rules, and an initial universe.
// Also returns an error, if one occurs.
func Parse(r io.Reader) ([]Rule, Universe, error) {
	scan := bufio.NewScanner(r)
	scan.Split(bufio.SplitFunc(splitter))

	var hasInput bool

	var lines []Rule
	var current Rule

	leftSide := true
	lineNum := 1

	for scan.Scan() {
		text := strings.TrimSpace(scan.Text())
		if strings.HasSuffix(text, "!") {
			text = text[:len(text)-1]
			hasInput = true
		}

		var err error
		if leftSide {
			err = current.Left.Parse(text)
		} else {
			err = current.Right.Parse(text)
		}
		if err != nil {
			if sErr, ok := err.(*SyntaxError); ok {
				sErr.line = lineNum
			}
			return nil, nil, err
		}

		// advance line number if we just scanned the rightward side
		if !leftSide {
			lines = append(lines, current)
		}
		if hasInput {
			break
		}
		if !leftSide {
			current = Rule{}
			lineNum++
		}
		leftSide = !leftSide
	}

	// scan for initial universe
	universe := Universe{}
	if hasInput && scan.Scan() {
		input := scan.Text()

		var rule Rule
		err := rule.Left.Parse(input)
		if err != nil {
			if sErr, ok := err.(*SyntaxError); ok {
				sErr.line = lineNum
			}
			return nil, Universe{}, err
		}

		universe = Universe(rule.Left)
	} else {
		universe = make(Universe)
		universe[SimpleAtom("_")] = 1
	}

	return lines, universe, nil
}
