package assembler

import (
	"fmt"
)

func Assemble(nodeList []Node) []uint8 {
	var pc uint16 = 0x0600

	var assembled []uint8

	for _, n := range nodeList {
		pc += uint16(n.size())
		if l, ok := n.(*labelNode); ok {
			(*l).address = pc
			// TODO: Index Labels
		}
	}

	fmt.Println(nodeList)

	for _, n := range nodeList {
		if i, ok := n.(*instrNode); ok {
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

	fmt.Println(Hexdump(assembled, 0, len(assembled)))

	return assembled
}

func Hexdump(byteArray []uint8, startPoint uint16, length int) string {
	var output string
	if len(byteArray) < int(startPoint)+length {
		return ""
	}

	for i := int(startPoint); i < length; i += 16 {
		output += fmt.Sprintf("%04x: ", i)
		for j := i; j < length && j-i < 16; j++ {
			output += fmt.Sprintf("%02x ", byteArray[j])
		}
		output += "\n"
	}
	return output
}
