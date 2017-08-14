package assembler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node interface {
	offset(uint16) uint8
	size() uint
}

type labelNode struct {
	content string
	address uint16
}

func (l labelNode) offset(dest uint16) uint8 {
	return uint8(l.address - dest)
}

func (l *labelNode) String() string {
	return fmt.Sprintf("<%v %#x>", l.content, l.address)
}

func (l *labelNode) size() uint {
	return 0
}

type instrNode struct {
	kind     string
	address  uint16
	location *labelNode
	mode     string
	opcode   uint8
}

func (i instrNode) offset(dest uint16) uint8 {
	return uint8(i.address - dest)
}

func (i *instrNode) String() string {
	if i.location == nil {
		return fmt.Sprintf("[$%02x]%v@<%v $%02x>", i.opcode, i.kind, i.mode, i.address)
	}
	return fmt.Sprintf("[$%02x]%v@%v", i.opcode, i.kind, i.location)
}

func (i *instrNode) size() uint {
	if i.mode == "imp" || i.mode == "acc" {
		return 1
	} else if i.mode == "ind" || strings.Contains(i.mode, "ab") {
		return 3
	} else {
		return 2
	}
}

const table = `,imm,acc,zp0,zpx,zpy,abs,abx,aby,ind,inx,iny,rel,imp
ADC,69,,65,75,,6d,7d,79,,61,72,,
AND,29,,25,35,,2d,3d,39,,21,31,,
ASL,,0a,6,16,,0e,1e,,,,,,
BCC,,,,,,,,,,,,90,
BCS,,,,,,,,,,,,b0,
BEQ,,,,,,,,,,,,f0,
BIT,,,24,,,2c,,,,,,,
BMI,,,,,,,,,,,,30,
BNE,,,,,,,,,,,,d0,
BPL,,,,,,,,,,,,10,
BRK,,,,,,,,,,,,,0
BVC,,,,,,,,,,,,50,
BVS,,,,,,,,,,,,70,
CLC,,,,,,,,,,,,,18
CLD,,,,,,,,,,,,,d8
CLI,,,,,,,,,,,,,58
CLV,,,,,,,,,,,,,b8
CMP,c9,,c5,d5,,cd,dd,d9,,c1,d1,,
CPX,e0,,e4,,,ec,,,,,,,
CPY,c0,,c4,,,cc,,,,,,,
DEC,,,c6,d6,,ce,de,,,,,,
DEX,,,,,,,,,,,,,ca
DEY,,,,,,,,,,,,,88
EOR,49,,45,55,,4d,5d,59,,41,51,,
INC,,,e6,f6,,ee,fe,,,,,,
INX,,,,,,,,,,,,,e8
INY,,,,,,,,,,,,,c8
JMP,,,,,,4c,,,6c,,,,
JSR,,,,,,20,,,,,,,
LDA,a9,,a5,b5,,ad,bd,b9,,a1,b1,,
LDX,a2,,a6,,b6,ae,,be,,,,,
LDY,a0,,a4,b4,,ac,bc,,,,,,
LSR,,4a,46,56,,4e,5e,,,,,,
NOP,,,,,,,,,,,,,ea
ORA,9,,5,15,,0d,1d,19,,1,11,,
PHA,,,,,,,,,,,,,48
PHP,,,,,,,,,,,,,8
PLA,,,,,,,,,,,,,68
PLP,,,,,,,,,,,,,28
ROL,,2a,26,36,,2e,3e,,,,,,
ROR,,6a,66,76,,6e,7e,,,,,,
RTI,,,,,,,,,,,,,40
RTS,,,,,,,,,,,,,60
SBC,e9,,e5,f5,,ed,fd,f9,,e1,f1,,
SEC,,,,,,,,,,,,,38
SED,,,,,,,,,,,,,f8
SEI,,,,,,,,,,,,,78
STA,,,85,95,,8d,9d,99,,81,91,,
STX,,,86,,96,8e,,,,,,,
STY,,,84,94,,8c,,,,,,,
TAX,,,,,,,,,,,,,aa
TAY,,,,,,,,,,,,,a8
TSX,,,,,,,,,,,,,ba
TXA,,,,,,,,,,,,,8a
TXS,,,,,,,,,,,,,9a
TYA,,,,,,,,,,,,,98
`

var lookup [][]string

func lookupTable(kind, mode string) (bool, uint8) {

	if len(lookup) == 0 {
		r := csv.NewReader(strings.NewReader(table))
		lookup, _ = r.ReadAll()
	}

	kind_index := -1

	for index, elem := range lookup {
		if elem[0] == kind {
			kind_index = index
			break
		}
	}

	if kind_index < 0 {
		return false, 0 // Did not find instruction
	}

	mode_index := -1

	for index, elem := range lookup[0] {
		if elem == mode {
			mode_index = index
			break
		}
	}

	if mode_index < 0 {
		return false, 0 // Did not find mode
	}

	if output, err := strconv.ParseUint(lookup[kind_index][mode_index], 16, 8); err == nil {
		return true, uint8(output)
	} else {
		return false, 0
	}

}

func NewInstrNodeFromAddr(kind string, addr uint16, mode string) (*instrNode, error) {
	if result, code := lookupTable(kind, mode); !result {
		msg := fmt.Sprintf("%v does not support addressing mode %v", kind, mode)
		return nil, errors.New(msg)
	} else {
		n := instrNode{kind: kind, address: addr, location: nil, mode: mode, opcode: code}
		return &n, nil
	}

}

func NewInstrNodeFromLabel(kind string, location *labelNode, mode string) (*instrNode, error) {
	if result, code := lookupTable(kind, mode); !result {
		msg := fmt.Sprintf("%v does not support addressing mode %v", kind, mode)
		return nil, errors.New(msg)
	} else {
		n := instrNode{kind: kind, address: 0, location: location, mode: mode, opcode: code}
		return &n, nil
	}
}
