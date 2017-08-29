package main

import (
	"flag"
	"fmt"
	"github.com/tusharjois/sfot/assembler"
	"github.com/tusharjois/sfot/simulator"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func assemble(str string) ([]byte, error) {
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

func run(st *simulator.State, debug bool) {
	isRunning := true

	for isRunning {
		isRunning = st.Step()
		if debug {
			fmt.Println(st)
			fmt.Println(st.HexdumpMemory(0, 0xff))
		}
	}
}

func helpRepl(isDebug bool) {
	if !isDebug {
		fmt.Println("load - load a file by name into sfot")
		fmt.Println("assemble - assemble loaded file")
		fmt.Println("run - run simulator on assembled program")
		fmt.Println("hexdump - display hexdump of assembled program")
		fmt.Println("disassemble - disassembly of assembled program")
	}
	fmt.Println("reset - reset program execution to original state")
	fmt.Println("debug - turn on debug mode")
	fmt.Println("step - step forward an instruction (debug only)")
	fmt.Println("jump - jump to program counter (debug only)")
	fmt.Println("print - print current processor state")
	fmt.Println("help - show this help message")
	if !isDebug {
		fmt.Println("exit - exit sfot")
	} else {
		fmt.Println("exit - exit debug mode")
	}
}

func debug(st *simulator.State) {
	fmt.Println("sfot debug mode active")
	isDebug := true

	var command string
	var subcommand string

	for isDebug {
		fmt.Print("sfot debug> ")
		fmt.Scanf("%s %x\n", &command, &subcommand)
		command = strings.ToLower(command)
		subcommand = strings.ToLower(subcommand)

		switch command {
		case "help":
			helpRepl(true)
		case "step":
			if !st.Step() {
				st
			}
		case "print":
			fmt.Println(st)
		case "jump":

		case "exit":
			return
		}
	}
}

func repl() {
	// commands := []string{"load", "assemble", "run", "reset", "hexdump", "disassemble", "debug", "step", "jump", "exit", "help", "print"}
	var command string
	var subcommand string

	fmt.Println("the sfot 6502 assembler and simulator")
	fmt.Println("type 'help' for a list of commands")

	var isRepl = true
	var debugFlag = false
	var currentState *simulator.State
	var fileData string
	var assembledProgram []byte

	for isRepl {
		fmt.Print("sfot> ")
		fmt.Scanf("%s %s\n", &command, &subcommand)
		command = strings.ToLower(command)
		subcommand = strings.ToLower(subcommand)
	}
}

func main() {

	// Command line flags
	assemble := flag.Bool("a", false, "run the assembler")
	simulate := flag.Bool("s", false, "run the simulator")
	debug := flag.Bool("b", false, "run the simulator in debug mode")
	disassemble := flag.Bool("d", false, "disassemble an assembled file")
	hexdump := flag.Bool("h", false, "hexdump an assembled file")
	// infile := flag.String("file", "", "provide an input file")  TODO
	// no_gfx := flag.Bool("no-gfx", false, "set to disable graphical display") TODO
	flag.Parse()

	var instream io.Reader = os.Stdin
	input, err := ioutil.ReadAll(instream)
	handleError(err)

	str := string(input)

	var program []byte

	if *assemble {
		program, err = runAssembler(str)
		handleError(err)
	} else {
		program = []byte(str)
	}

	if *simulate {
		runSimulator(program, false) // TODO
	} else if *debug {
		runSimulator(program, true) // TODO
	} else if *disassemble {
		runSimulator(program, true) // TODO
	} else if *hexdump {
		runSimulator(program, true) // TODO
	}

	repl()

}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[\x1b[31merror\x1b[0m] %v\n", err)
		os.Exit(1)
	}
}
