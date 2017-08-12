package assembler

import (
	"errors"
)

type node interface {
	offset(uint16) uint8
}

type labelNode struct {
	content string
	address uint16
}

func (l labelNode) offset(dest uint16) uint8 {
	return from - l.address
}

type instrNode struct {
	kind     string
	address  uint16
	location *label
	mode     string
	opcode   uint8
}

func (i instrNode) offset(dest uint16) uint8 {
	return from - i.address
}

func lookupTable(kind, mode string) bool {
	imm := []string{"ADC", "AND", "CMP", "CPX", "CPY", "EOR", "LDA", "LDX",
		"LDY", "ORA", "SBC"}
	acc := []string{"ASL", "LSR", "ROL", "ROR"}
	zp0 := []string{"ADC", "AND", "ASL", "BIT", "CMP", "CPX", "CPY", "DEC",
		"EOR", "INC", "LDA", "LDX", "LDY", "LSR", "ORA", "ROL", "ROR", "SBC",
		"STA", "STX", "STY"}
	zpx := []string{"ADC", "AND", "ASL", "CMP", "DEC", "EOR", "INC", "LDA",
		"LDY", "LSR", "ORA", "ROL", "ROR", "SBC", "STA", "STY"}
	zpy := []string{"LDX", "STX"}
	abs := []string{"ADC", "AND", "ASL", "BIT", "CMP", "CPX", "CPY", "DEC",
		"EOR", "INC", "JMP", "JSR", "LDA", "LDX", "LDY", "LSR", "ORA", "ROL",
		"ROR", "SBC", "STA", "STX", "STY"}
	abx := []string{"ADC", "AND", "ASL", "CMP", "DEC", "EOR", "INC", "LDA",
		"LDY", "LSR", "ORA", "ROL", "ROR", "SBC", "STA"}
	aby := []string{"ADC", "AND", "CMP", "EOR", "LDA", "LDX", "ORA", "SBC",
		"STA"}
	ind := []string{"JMP"}
	inx := []string{"ADC", "AND", "CMP", "EOR", "LDA", "ORA", "SBC", "STA"}
	iny := []string{"ADC", "AND", "CMP", "EOR", "LDA", "ORA", "SBC", "STA"}
	rel := []string{"BPL", "BMI", "BVC", "BCC", "BCS", "BNE", "BEQ"}
	imp := []string{"BRK", "CLC", "SEC", "CLI", "SEI", "CLV", "CLD", "SED",
		"NOP", "TAX", "TXA", "DEX", "INX", "TAY", "TYA", "DEY", "INX", "RTI",
		"RTS", "TXS", "TSX", "PHA", "PLA", "PHP", "PLP"}

	var vaild []string

	switch mode {
	case "imm":
		valid = imm
	case "acc":
		valid = acc
	case "zp0":
		valid = zp0
	case "zpx":
		valid = zpx
	case "zpy":
		valid = zpy
	case "abs":
		valid = abs
	case "abx":
		valid = abx
	case "aby":
		valid = aby
	case "ind":
		valid = ind
	case "inx":
		valid = inx
	case "iny":
		valid = iny
	case "rel":
		valid = rel
	case "imp":
		valid = imp
	default:
		return false
	}

	for _, elem := range valid {
		if elem == mode {
			return true
		}
	}

	return false
}

func NewInstrNodeFromAddr(kind string, addr uint16, mode string) (instrNode, error) {

}

func NewInstrNodeFromLabel(kind string, location *label, mode string) (instrNode, error) {

}
