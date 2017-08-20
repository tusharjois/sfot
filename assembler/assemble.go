package assembler

import (
	"errors"
	"fmt"
)

func Assemble(nodeList []Node) ([]uint8, error) {
	// Starts assembly at PC=$0800
	var pc uint16 = 0x0800

	assembled := []uint8{'s', 'f', 'o', 't'} // Magic number
	labelIndex := make(map[string]bool)

	for _, n := range nodeList {
		pc += uint16(n.size())
		if l, ok := n.(*labelNode); ok {
			(*l).address = pc
			labelIndex[l.content] = true
		}
	}

	pc = 0x0800

	for _, n := range nodeList {
		if i, ok := n.(*instrNode); ok {
			pc += uint16(i.size()) // Increment PC

			if i.location != nil {
				if _, prs := labelIndex[i.location.content]; !prs {
					return assembled,
						errors.New(fmt.Sprintf("could not find label %v",
							i.location.content))
				}

				if i.mode == "rel" {
					i.address = uint16(i.location.offset(pc))
				} else {
					i.address = i.location.address
				}
			}

			switch i.size() {
			case 1:
				assembled = append(assembled, i.opcode)
			case 2:
				assembled = append(assembled, i.opcode)
				assembled = append(assembled, uint8(i.address))
			case 3:
				assembled = append(assembled, i.opcode)
				assembled = append(assembled, uint8(i.address&0xff))      // little-end
				assembled = append(assembled, uint8((i.address>>8)&0xff)) // big-end
			}

		}
	}

	return assembled, nil
}

func Hexdump(byteArray []uint8) string {
	var output string

	for i := 0; i < len(byteArray); i += 16 {
		output += fmt.Sprintf("%04x: ", i)
		for j := i; j < len(byteArray) && j-i < 16; j++ {
			output += fmt.Sprintf("%02x ", byteArray[j])
		}
		output += "\n"
	}
	return output
}
