package main

import (
	"github.com/tusharjois/sfot/assembler"
	"github.com/tusharjois/sfot/simulator"
	"testing"
)

func prepare(str string) ([]byte, error) {
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

func TestAddressingModes(t *testing.T) {
	const value = 0xfa
	addr_tables := []struct {
		str  string
		name string
	}{
		{"LDA #$fa", "Immediate"},
		{"LDX #$fa\nSTX $aaaa\nLDA $aaaa", "Absolute"},
		{"LDX #$fa\nSTX $aa\nLDA $aa", "Zero page"},
		{"LDY #$fa\nINX\nSTY $ab\nLDA $aa,X", "Zero page,X"},
		{"LDX #$fa\nINY\nSTX $ab,Y\nLDA $ac", "Zero page,Y"},
		{"LDY #$fa\nINX\nSTY $aaab\nLDA $aaaa,X", "Absolute,X"},
		{"LDX #$fa\nINY\nSTX $aaab\nLDA $aaaa,Y", "Absolute,Y"},
		{"LDA #$fa\nLSR", "Implicit"},
		// {"LDX #$fa\nJMP post\nINX\npost:\nTXA", "Indirect"}, TODO
		{"LDX #$fa\nSEC\nBCS post\nINX\npost:\nTXA", "Relative"},
		{"LDX #$01\nLDA #$05\nSTA $01\nLDA #$09\nSTA $02\nLDY #$fa\nSTY $0905\nLDA ($00,X)", "Indexed indirect"},
		{"LDY #$01\nLDA #$03\nSTA $01\nLDA #$09\nSTA $02\nLDX #$fa\nSTX $0904\nLDA ($01),Y", "Indirect indexed"},
	}

	for _, table := range addr_tables {
		t.Log(table.name)
		program, err := prepare(table.str)
		t.Log(simulator.Disassemble(program))
		st := simulator.NewState(program)
		t.Log()
		if err != nil {
			t.Error(err)
		}

		for st.Step() {
		}

		acc := st.GetByteRegisters()["A"]

		if acc != value {
			t.Errorf("Incorrect value in accumulator for %s mode. Want $fa, got $%x", table.name, acc)
		}
	}

}

func TestAddressingEdgeCases(t *testing.T) {}

func TestProcessorFlagInstructions(t *testing.T) {}

func TestAccumulatorInstructions(t *testing.T) {}
