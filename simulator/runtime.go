package simulator

import (
	"fmt"
	"strings"
)

func (st *State) Step() bool {
	opcode := st.memory[st.ProgramCounter]
	st.ProgramCounter++

	info := strings.Split(opcodeMatrix[opcode], "-")
	// info[0] is the name of the opcode, info[1] is the addressing mode

	var location uint16
	var immediate byte
	var immediateFlag bool = false

	// TODO: Check Panic: program counter overflow
	// TODO: Check Panic: index out of bounds for opcode
	switch info[1] {
	case "imm":
		immediate = st.memory[st.ProgramCounter]
		st.ProgramCounter++
		immediateFlag = true
	case "acc":
		immediateFlag = true
	case "zp0":
		location = uint16(st.memory[st.ProgramCounter])
		st.ProgramCounter++
	case "zpx":
		location = uint16(st.memory[st.ProgramCounter]) + uint16(st.indexX)
		st.ProgramCounter++
		if location > 0xff {
			location &= 0xfe
		}
	case "zpy":
		location = uint16(st.memory[st.ProgramCounter]) + uint16(st.indexY)
		st.ProgramCounter++
		if location > 0xff {
			location &= 0xfe
		}
	case "abs":
		location = uint16(st.memory[st.ProgramCounter])
		st.ProgramCounter++
		location += uint16(uint32(st.memory[st.ProgramCounter]) << 8)
		st.ProgramCounter++
	case "abx":
		location = uint16(st.memory[st.ProgramCounter])
		st.ProgramCounter++
		location += uint16(st.memory[st.ProgramCounter]) << 8
		st.ProgramCounter++
		tempLoc := uint32(st.indexX) + uint32(location)
		if tempLoc > 0xffff {
			tempLoc -= 0xffff
		}
		location = uint16(tempLoc)
	case "aby":
		location = uint16(st.memory[st.ProgramCounter])
		st.ProgramCounter++
		location += uint16(st.memory[st.ProgramCounter]) << 8
		st.ProgramCounter++
		tempLoc := uint32(st.indexY) + uint32(location)
		if tempLoc > 0xffff {
			tempLoc -= 0xffff
		}
		location = uint16(tempLoc)
	case "ind":
		indirectLoc := uint16(st.memory[st.ProgramCounter])
		st.ProgramCounter++
		location = uint16(st.memory[indirectLoc])
		indirectLoc++
		if indirectLoc > 0xff {
			indirectLoc &= 0xfe
		}
		location += uint16(st.memory[indirectLoc]) << 8
	case "inx":
		indirectLoc := uint16(st.memory[st.ProgramCounter]) + uint16(st.indexX)
		st.ProgramCounter++
		if indirectLoc > 0xff {
			indirectLoc &= 0xfe
		}
		location = uint16(st.memory[indirectLoc])
		indirectLoc++
		if indirectLoc > 0xff {
			indirectLoc &= 0xfe
		}
		location += uint16(st.memory[indirectLoc]) << 8
	case "iny":
		indirectLoc := st.memory[st.ProgramCounter]
		st.ProgramCounter++
		location = uint16(st.memory[indirectLoc])
		indirectLoc++
		if indirectLoc > 0xff {
			indirectLoc &= 0xfe
		}
		location += uint16(st.memory[indirectLoc]) << 8
		tempLoc := uint32(st.indexY) + uint32(location)
		if tempLoc > 0xffff {
			tempLoc -= 0xffff
		}
		location = uint16(tempLoc)
	case "rel":
		offset_rel := int32(int8(st.memory[st.ProgramCounter])) + int32(st.ProgramCounter)
		fmt.Printf("%04x\n", offset_rel)
		if offset_rel > 0xffff {
			offset_rel -= 0xffff
		} else if offset_rel < 0x0000 {
			offset_rel += 0xffff
		}
		location = uint16(offset_rel)
	case "imp":
	}

	st.handleInstruction(info[0], location, immediate, immediateFlag)
	return st.breakFlag == 0
}

func (st *State) handleInstruction(instr string, location uint16, immediate byte, flag bool) {

	switch instr {
	case "ADC":
		st.adc(location, immediate, flag)
	case "AND":
		st.and(location, immediate, flag)
	case "ASL":
		st.asl()
	case "BCC":
		st.bcc(location)
	case "BCS":
		st.bcs(location)
	case "BEQ":
		st.beq(location)
	case "BIT":
		st.bit()
	case "BMI":
		st.bmi(location)
	case "BNE":
		st.bne(location)
	case "BPL":
		st.bpl(location)
	case "BRK":
		st.brk()
	case "BVC":
		st.bvc(location)
	case "BVS":
		st.bvs(location)
	case "CLC":
		st.clc()
	case "CLD":
		st.cld()
	case "CLI":
		st.cli()
	case "CLV":
		st.clv()
	case "CMP":
		st.cmp()
	case "CPX":
		st.cpx()
	case "CPY":
		st.cpy()
	case "DEC":
		st.dec(location)
	case "DEX":
		st.dex()
	case "DEY":
		st.dey()
	case "EOR":
		st.eor()
	case "INC":
		st.inc(location)
	case "INX":
		st.inx()
	case "INY":
		st.iny()
	case "JMP":
		st.jmp()
	case "JSR":
		st.jsr()
	case "LDA":
		st.lda(location, immediate, flag)
	case "LDX":
		st.ldx(location, immediate, flag)
	case "LDY":
		st.ldy(location, immediate, flag)
	case "LSR":
		st.lsr()
	case "NOP":
		st.nop()
	case "ORA":
		st.ora()
	case "PHA":
		st.pha()
	case "PHP":
		st.php()
	case "PLA":
		st.pla()
	case "PLP":
		st.plp()
	case "ROL":
		st.rol()
	case "ROR":
		st.ror()
	case "RTI":
		st.rti()
	case "RTS":
		st.rts()
	case "SBC":
		st.sbc()
	case "SEC":
		st.sec()
	case "SED":
		st.sed()
	case "SEI":
		st.sei()
	case "STA":
		st.sta(location)
	case "STX":
		st.stx(location)
	case "STY":
		st.sty(location)
	case "TAX":
		st.tax()
	case "TAY":
		st.tay()
	case "TSX":
		st.tsx()
	case "TXA":
		st.txa()
	case "TXS":
		st.txs()
	case "TYA":
		st.tya()
	}

}
