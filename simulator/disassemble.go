package simulator

import (
	"fmt"
	"strings"
)

func Disassemble(assembled []byte) string {
	if len(assembled) < 4 || (assembled[0] != 's' && assembled[1] != 'f' && assembled[2] != 'o' && assembled[3] != 't') {
		return "Invalid sfot binary (incorrect magic number)\n"
	}

	output := "Address    Hexdump    Disassembled\n"
	output += "========   ========   =============\n"

	address := 0x800 // Start of program counter

	index := 4
	for index < len(assembled) {
		opcode := assembled[index]
		info := strings.Split(opcodeMatrix[opcode], "-")
		// info[0] is the name of the opcode, info[1] is the addressing mode

		hexdump := ""
		instr_part := info[0]
		loc_part := ""
		jump_amt := 0

		switch info[1] {
		case "imm":
			loc_part = fmt.Sprintf("$#%02x", assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "acc":
			loc_part = "A"
			jump_amt++
		case "zp0":
			loc_part = fmt.Sprintf("$%02x", assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "zpx":
			loc_part = fmt.Sprintf("$%02x,X", assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "zpy":
			loc_part = fmt.Sprintf("$%02x,Y", assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "abs":
			loc_part = fmt.Sprintf("$%02x%02x", assembled[index+2], assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x %02x", assembled[index], assembled[index+1], assembled[index+2])
			jump_amt += 3
		case "abx":
			loc_part = fmt.Sprintf("$%02x%02x,X", assembled[index+2], assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x %02x", assembled[index], assembled[index+1], assembled[index+2])
			jump_amt += 3
		case "aby":
			loc_part = fmt.Sprintf("$%02x%02x,Y", assembled[index+2], assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x %02x", assembled[index], assembled[index+1], assembled[index+2])
			jump_amt += 3
		case "ind":
			loc_part = fmt.Sprintf("($%02x%02x)", assembled[index+2], assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x %02x", assembled[index], assembled[index+1], assembled[index+2])
			jump_amt += 3
		case "inx":
			loc_part = fmt.Sprintf("($%02x,X)", assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "iny":
			loc_part = fmt.Sprintf("($%02x),Y", assembled[index+1])
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "rel":
			loc_part = fmt.Sprintf("$%04x", int(int8((assembled[index+1])))+address+2)
			hexdump = fmt.Sprintf("%02x %02x", assembled[index], assembled[index+1])
			jump_amt += 2
		case "imp":
			hexdump = fmt.Sprintf("%02x", assembled[index])
			jump_amt++
		}

		output += fmt.Sprintf("$%04x      %-8v   %v %v\n", address, hexdump, instr_part, loc_part)
		address += jump_amt
		index += jump_amt
	}

	return output
}
