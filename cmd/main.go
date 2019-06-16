package main

import (
	"fmt"
	"strings"

	"github.com/deanveloper/alchemist"
)

func main() {
	test := `
	_ -> Out_"Enter how many numbers you wanna see:"+In_loop+b+setNext+Out_""+Out_"Fibonacci:"+Out_a+Out_b

	loop+a+setNext -> loop+next+setNext
	loop+b+setNext -> loop+next+setNext+saveB
	loop+0a+0b+setNext -> Out_next+setA

	loop+setA+saveB -> loop+setA+a
	loop+setA+0saveB -> loop+setB

	loop+setB+next -> loop+setB+b
	loop+setB+0next -> loop+setNext
`
	rules, universe, err := alchemist.Parse(strings.NewReader(test))
	if err != nil {
		fmt.Println("err", err)
	}

	universe.Run(rules)
}
