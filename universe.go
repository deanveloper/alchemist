package alchemist

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

// Universe represents the state of the program, which is
// the number of each atom.
type Universe map[SimpleAtom]int

// Step incremements the universe by a step using the given rules.
//
// Returns whether the step was successful.
func (u Universe) Step(rules []Rule) bool {

	var applicable []Rule

	for _, rule := range rules {
		reqsMet := true
		for atom, num := range rule.Left {
			if num == 0 && u[atom] != 0 {
				reqsMet = false
				break
			}
			if u[atom] < num {
				reqsMet = false
				break
			}
		}

		if reqsMet {
			applicable = append(applicable, rule)
		}
	}

	if len(applicable) == 0 {
		return false
	}

	r := random.Intn(len(applicable))
	rule := applicable[r]

	for a, n := range rule.Left {
		u[a] -= n
		if u[a] == 0 {
			delete(u, a)
		}
	}

	for _, a := range rule.Right {
		a.Run(u)
	}

	return true
}

// Run will continuously call u.Step(rules) until it returns false.
func (u Universe) Run(rules []Rule) {
	for u.Step(rules) {
	}
}
