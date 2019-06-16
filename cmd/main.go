package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deanveloper/alchemist"
)

var (
	inPath       = flag.String("i", "[file]", "input: the file to read in from")
	rules        = flag.String("r", "[rules]", "rules: an inline way to declare rules, overrides -i")
	universeFlag = flag.String("u", "[universe]", "universe: the initial atoms in the universe")
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s [flags]\n", os.Args[0])
		fmt.Printf("Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	rules, universe := parseRules()

	// inline universe overrides universe provided by file
	if *universeFlag != "[universe]" {
		var univ alchemist.LHSRule
		err := univ.Parse(*universeFlag)
		if err != nil {
			log.Fatalln("error parsing -u:", err)
		}
		universe = alchemist.Universe(univ)
	}

	universe.Run(rules)
}

func parseRules() ([]alchemist.Rule, alchemist.Universe) {
	if *rules != "[rules]" {
		rules, universe, err := alchemist.Parse(strings.NewReader(*rules))
		if err != nil {
			log.Fatalf("error parsing -r: %v", err)
		}
		return rules, universe
	}

	if *inPath != "[file]" {
		f, err := os.Open(*inPath)
		if err != nil {
			log.Fatalf("error opening %q: %v", *inPath, err)
		}

		rules, universe, err := alchemist.Parse(f)
		if err != nil {
			log.Fatalf("error parsing %q: %v", *inPath, err)
		}

		return rules, universe
	}

	log.Fatalln("required: either -i or -r")
	return nil, nil
}
