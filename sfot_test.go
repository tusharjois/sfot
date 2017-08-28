package test

import (
	"github.com/tusharjois/sfot/assembler"
	"github.com/tusharjois/sfot/simulator"
	"testing"
)

func prepare(str string) (*State, error) {
	var program []byte

	str, _, err := assembler.Preprocess(str)
	if err != nil {
		return nil, err
	}

	tz, err := assembler.NewTokenizer(&str)
	if err != nil {
		return nil, err
	}

	p, err := assembler.Parse(tz)
	if err != nil {
		return nil, err
	}

	program, err = assembler.Assemble(p)
	if err != nil {
		return nil, err
	}

	st := simulator.NewState(program)

	return st, nil
}

func TestAddressingModes(t *testing.T) {
	const value = 0xfa
	addr_table := []struct {
		instr string
		name  string
	}{
		{"LDA #$fa", "Immediate"},
		{"LDX #$fa\nSTX $aaaa\nLDA $aaaa", "Absolute"},
		{"LDX #$fa\nSTX $aa\nLDA $aa", "Zero page"},
		{"LDY #$fa\nINX\nSTY $ab\nLDA $aa,X", "Zero page,X"},
		{"LDX #$fa\nINY\nSTX $ab\nLDA $aa,Y", "Zero page,Y"},
		{"LDY #$fa\nINX\nSTY $aaab\nLDA $aaaa,X", "Absolute,X"},
		{"LDX #$fa\nINY\nSTX $aaab\nLDA $aaaa,Y", "Absolute,Y"},
		{"LDA #$fa\nLSL", "Implicit"},
		{"LDX #$fa\nJMP post\nINX\npost:\nTXA", "Indirect"},
		{"LDX #$fa\nSEC\nBCS post\nINX\npost:\nTXA", "Relative"},
		{"LDA #$fa\nSTA $aa\nLDA $aa", "Indexed indirect"},
		{"LDA #$fa\nSTA $aa\nLDA $aa", "Indirect indexed"},
	}

	st, err := prepare(str)
	if err != nil {
		t.Error(err)
	}

	if !st.Step() {
		t.Error("Simulator step failed.")
	}

	acc := st.GetByteRegisters()["A"]

	if acc != value {
		t.Errorf("Incorrect value in accumulator. Want $fa, got $%x", acc)
	}
}

func TestAddressingEdgeCases(t *testing.T) {}

func TestProcessorFlagInstructions(t *testing.T) {}

func TestAccumulatorInstructions(t *testing.T) {}
