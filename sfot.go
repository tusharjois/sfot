package main

import (
	"fmt"
	"github.com/tusharjois/sfot/assembler"
	"github.com/tusharjois/sfot/simulator"
	"io/ioutil"
	"os"
)

func runAssembler(str string) ([]byte, error) {
	var program []byte

	str, _, err := assembler.Preprocess(str)
	if err != nil {
		return program, err
	}

	tz, err := assembler.NewTokenizer(&str)
	if err != nil {
		return program, err
	}

	p, err := assembler.Parse(tz)
	if err != nil {
		return program, err
	}

	program, err = assembler.Assemble(p)
	if err != nil {
		return program, err
	}

	return program, nil
}

func runSimulator(program []byte, debug bool) {
	isRunning := true
	st := simulator.NewState(program)

	for isRunning {
		isRunning = st.Step()
		if debug {
			fmt.Println(st)
			fmt.Println(st.HexdumpMemory(0, 0xff))
		}
	}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	handleError(err)

	str := string(input)

	program, err := runAssembler(str)
	handleError(err)

	runSimulator(program, true) // TODO
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[\x1b[31merror\x1b[0m] %v\n", err)
		os.Exit(1)
	}
}
