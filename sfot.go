package main

import (
	"fmt"
	"github.com/tusharjois/sfot/assembler"
	"github.com/tusharjois/sfot/simulator"
	"io/ioutil"
	"os"
)

func main() {

	input, err := ioutil.ReadAll(os.Stdin)

	str := string(input)

	str, _, err = assembler.Preprocess(str)
	fmt.Println(str)
	handleError(err)

	tz, err := assembler.NewTokenizer(&str)
	handleError(err)

	p, err := assembler.Parse(tz)
	handleError(err)

	program, err := assembler.Assemble(p)
	handleError(err)

	fmt.Println(assembler.Hexdump(program))
	fmt.Println(simulator.Disassemble(program))
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[\x1b[31merror\x1b[0m] %v\n", err)
		os.Exit(1)
	}
}
