package simulator

import "fmt"

type State struct {
	// Registers
	indexX         byte
	indexY         byte
	accumulator    byte
	stackPointer   byte
	ProgramCounter uint16

	// Flags - 0 or 1
	carry     byte
	zero      byte
	interrupt byte
	decimal   byte
	breakFlag byte
	overflow  byte
	negative  byte

	// Memory
	memory [0xffff]byte
}

func (st *State) String() string {
	registers := fmt.Sprintf("A=$%02x X=$%02x Y=$%02x\n SP=$%02x PC=$%04x\n",
		st.accumulator, st.indexX, st.indexY, st.stackPointer,
		st.ProgramCounter)

	flags := fmt.Sprintf("NV-BDIZC\n%v%v1%v%v%v%v%v", st.negative, st.overflow,
		st.breakFlag, st.decimal, st.interrupt, st.zero, st.carry)

	return fmt.Sprintf("%v\n%v", registers, flags)
}

func (st *State) GetProcessorFlags() map[string]bool {
	m := make(map[string]bool)

	m["C"] = st.carry == 1
	m["Z"] = st.zero == 1
	m["I"] = st.interrupt == 1
	m["D"] = st.decimal == 1
	m["B"] = st.breakFlag == 1
	m["V"] = st.overflow == 1
	m["N"] = st.negative == 1

	return m
}

func (st *State) GetByteRegisters() map[string]byte {
	m := make(map[string]byte)

	m["X"] = st.indexX
	m["A"] = st.accumulator
	m["Y"] = st.indexY
	m["SP"] = st.stackPointer

	return m
}

func (st *State) HexdumpMemory(startPoint, length uint16) string {
	var output string

	if startPoint+length >= 0xffff {
		length = 0xffff
	}

	for i := startPoint; i < length; i += 16 {
		output += fmt.Sprintf("%04x: ", i+0x0800)
		for j := i; j < length && j-i < 16; j++ {
			output += fmt.Sprintf("%02x ", st.memory[j])
		}
		output += "\n"
	}
	return output
}

func NewState(program []byte) *State {
	if len(program)+0x0800 >= 0xffff {
		return nil
	}

	st := new(State)

	st.Reset()

	for _, element := range program {
		st.memory[st.ProgramCounter] = element
		st.ProgramCounter++
	}

	st.ProgramCounter = 0x0800

	return st
}

func (st *State) Reset() {
	st.accumulator = 0x0
	st.indexX = 0x0
	st.indexY = 0x0
	st.stackPointer = 0xff // Page 1 of memory, actual location 0x100 + this
	st.ProgramCounter = 0x0800

	st.carry = 0
	st.zero = 0
	st.interrupt = 0
	st.decimal = 0
	st.breakFlag = 1
	st.overflow = 0
	st.negative = 0
}
