package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/tusharjois/sfot/assembler"
	"github.com/tusharjois/sfot/simulator"
	"io/ioutil"
	"os"
	"strings"
)

func assemble(str string) ([]byte, error) {
	var program []byte

	if str == "" {
		return program, errors.New("program is empty")
	}

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

func run(st *simulator.State) {
	for st.Step() {
	}

	fmt.Println("program stopped execution (pc=$%04x)", st.ProgramCounter)
}

func helpRepl(isDebug bool) {
	if !isDebug {
		fmt.Println("load - load a file by name into sfot")
		fmt.Println("assemble - assemble loaded file")
		fmt.Println("run - run simulator on assembled program")
		fmt.Println("hexdump - display hexdump of assembled program")
		fmt.Println("disassemble - disassembly of assembled program")
	}
	fmt.Println("debug - toggle debug mode")
	fmt.Println("reset - reset program execution to original state")
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
	var command string
	var subcommand string

	for true {
		fmt.Print("sfot debug> ")
		fmt.Scanf("%s %s\n", &command, &subcommand)
		command = strings.ToLower(command)
		subcommand = strings.ToLower(subcommand)

		switch command {
		case "help":
			helpRepl(true)
		case "step":
			if !st.Step() {
				fmt.Println("program stopped execution (pc=$%04x)", st.ProgramCounter)
			}
		case "print":
			fmt.Println(st)
		case "reset":
			st.Reset()
		case "jump":
			// TODO: implement
		case "exit":
			return
		}
	}
}

func repl(str string) {
	// commands := []string{"load", "assemble", "run", "reset", "hexdump", "disassemble", "debug", "step", "jump", "exit", "help", "print"}
	var command string
	var subcommand string

	fmt.Println("the sfot 6502 assembler and simulator")
	fmt.Println("type 'help' for a list of commands\n")

	var currentState *simulator.State
	var fileData string = str

	for true {
		fmt.Print("sfot> ")
		fmt.Scanf("%s %s\n", &command, &subcommand)
		command = strings.ToLower(command)
		subcommand = strings.ToLower(subcommand)
		switch command {
		case "assemble":
			assembledProgram, err := assemble(fileData)
			handleError(err, true)
			currentState = simulator.NewState(assembledProgram)
		case "run":
			if currentState == nil {
				fmt.Println("no assembled program found")
			} else {
				run(currentState)
			}
		case "load":
			input, err := ioutil.ReadFile(subcommand)
			handleError(err, true)
			fileData = string(input)
			fmt.Printf("loaded file %q\n", subcommand)
		case "debug":
			if currentState == nil {
				fmt.Println("no assembled program found")
			} else {
				debug(currentState)
			}
		case "help":
			helpRepl(false)
		case "exit":
			return
		default:
			fmt.Println("invalid command")
		}
	}
}

func main() {

	// Command line flags
	assembly := flag.Bool("a", false, "run the assembler")
	simulate := flag.Bool("s", false, "run the simulator")
	debugMode := flag.Bool("b", false, "run the simulator in debug mode")
	disassembly := flag.Bool("d", false, "disassemble an assembled file")
	hexdumpProgram := flag.Bool("h", false, "hexdump an assembled file")
	infile := flag.String("file", "", "provide an input file")
	// no_gfx := flag.Bool("no-gfx", false, "set to disable graphical display") TODO
	flag.Parse()

	var str string

	if *infile == "" {
		input, err := ioutil.ReadAll(os.Stdin)
		handleError(err, false)
		str = string(input)
		// TODO repl placement
	} else {
		input, err := ioutil.ReadFile(*infile)
		handleError(err, false)
		str = string(input)
	}

	var program []byte

	if *assembly {
		var err error
		program, err = assemble(str)
		handleError(err, false)
	} else {
		program = []byte(str)
	}

	st := simulator.NewState(program)

	if st == nil {
		repl("")
	} else {
		if *simulate {
			run(st)
		} else if *debugMode {
			debug(st)
		} else if *disassembly {
			debug(st) // TODO
		} else if *hexdumpProgram {
			debug(st) // TODO
		} else {
			repl(str)
		}
	}

}

func handleError(err error, interactive bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[\x1b[31merror\x1b[0m] %v\n", err)
		if !interactive {
			os.Exit(1)
		}
	}
}
