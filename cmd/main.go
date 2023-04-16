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
	inPath       = flag.String("i", "", "input file to read rules and universe from")
	rules        = flag.String("r", "", "rules to execute with, overrides -i")
	universeFlag = flag.String("u", "", "universe of atoms to start with")
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s [flags]\n", os.Args[0])
		fmt.Printf("Flags:\n")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("To learn the language, read here: https://github.com/bforte/Alchemist/blob/master/README.md")
	}

	flag.Parse()

	rules, universe := parseRules()

	// inline universe overrides universe provided by file
	if *universeFlag != "" {
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
	if *rules != "" {
		rules, universe, err := alchemist.Parse(strings.NewReader(*rules))
		if err != nil {
			log.Fatalf("error parsing -r: %v", err)
		}
		return rules, universe
	}

	if *inPath != "" {
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
