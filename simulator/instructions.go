package simulator

import ()

func (st *State) adc(loc uint16, imm byte, flag bool) {
	var addition uint16
	if flag {
		addition = uint16(st.accumulator) + uint16(imm)
	} else {
		addition = uint16(st.accumulator) + uint16(st.memory[loc])
	}

	bitSeven := (addition >> 7) & 1

	st.negative = uint8(bitSeven)

	if addition > 0xff {
		st.carry = 1
	} else {
		st.overflow = 0
	}

	if addition == 0 {
		st.zero = 1
	} else {
		st.zero = 0
	}

	if int16(addition) > 127 || int16(addition) < -128 {
		st.overflow = 1
	} else {
		st.overflow = 0
	}

	st.accumulator = byte(addition & 255)
}

func (st *State) and(loc uint16, imm byte, flag bool) {
	var value byte
	if flag {
		value = imm
	} else {
		value = st.memory[loc]
	}

	st.accumulator = st.accumulator & value

	st.negative = (st.accumulator >> 7) & 1

	if st.accumulator == 0 {
		st.zero = 1
	} else {
		st.zero = 0
	}
}

func (st *State) asl() {
	// TODO
}

func (st *State) bcc(loc uint16) {
	if st.carry == 0 {
		st.ProgramCounter = loc
	}
}

func (st *State) bcs(loc uint16) {
	if st.carry == 1 {
		st.ProgramCounter = loc
	}
}

func (st *State) beq(loc uint16) {
	if st.zero == 1 {
		st.ProgramCounter = loc
	}
}

func (st *State) bit() {
	// TODO
}

func (st *State) bmi(loc uint16) {
	if st.negative == 1 {
		st.ProgramCounter = loc
	}
}

func (st *State) bne(loc uint16) {
	if st.zero == 0 {
		st.ProgramCounter = loc
	}
}

func (st *State) bpl(loc uint16) {
	if st.negative == 0 {
		st.ProgramCounter = loc
	}
}

func (st *State) brk() {
	st.breakFlag = 1
}

func (st *State) bvc(loc uint16) {
	if st.overflow == 0 {
		st.ProgramCounter = loc
	}
}

func (st *State) bvs(loc uint16) {
	if st.overflow == 1 {
		st.ProgramCounter = loc
	}
}

func (st *State) clc() {
	st.carry = 0
}

func (st *State) cld() {
	st.decimal = 0
}

func (st *State) cli() {
	st.interrupt = 0
}

func (st *State) clv() {
	st.overflow = 0
}

func (st *State) cmp() {
	// TODO
}

func (st *State) cpx() {
	// TODO
}

func (st *State) cpy() {
	// TODO
}

func (st *State) dec(loc uint16) {
	value := uint16(st.memory[loc])
	value--
	value = value & 255
	st.memory[loc] = byte(value)

	st.negative = byte((value >> 7) & 1)

	if value == 0 {
		st.zero = 1
	} else {
		st.zero = 1
	}

}

func (st *State) dex() {
	value := st.indexX
	value--
	value = value & 255
	st.indexX = value

	st.negative = (value >> 7) & 1

	if value == 0 {
		st.zero = 1
	} else {
		st.zero = 1
	}
}

func (st *State) dey() {
	value := st.indexY
	value--
	value = value & 255
	st.indexY = value

	st.negative = (value >> 7) & 1

	if value == 0 {
		st.zero = 1
	} else {
		st.zero = 1
	}
}

func (st *State) eor() {
	// TODO
}

func (st *State) inc(loc uint16) {
	value := uint16(st.memory[loc])
	value++
	value = value & 255
	st.memory[loc] = byte(value)

	st.negative = byte((value >> 7) & 1)

	if value == 0 {
		st.zero = 1
	} else {
		st.zero = 1
	}

}

func (st *State) inx() {
	value := st.indexX
	value++
	value = value & 255
	st.indexX = value

	st.negative = (value >> 7) & 1

	if value == 0 {
		st.zero = 1
	} else {
		st.zero = 1
	}
}

func (st *State) iny() {
	value := st.indexY
	value++
	value = value & 255
	st.indexY = value

	st.negative = (value >> 7) & 1

	if value == 0 {
		st.zero = 1
	} else {
		st.zero = 1
	}
}

func (st *State) jmp(loc uint16) {
	st.ProgramCounter = loc
}

func (st *State) jsr(loc uint16) {
	st.stackPush(byte(st.ProgramCounter >> 8))   // Push hi byte first
	st.stackPush(byte(st.ProgramCounter & 0xfe)) // Push lo byte next
	st.ProgramCounter = loc
}

func (st *State) lda(loc uint16, imm byte, flag bool) {
	if flag {
		st.accumulator = imm
	} else {
		st.accumulator = st.memory[loc]
	}
}

func (st *State) ldx(loc uint16, imm byte, flag bool) {
	if flag {
		st.indexX = imm
	} else {
		st.indexX = st.memory[loc]
	}
}

func (st *State) ldy(loc uint16, imm byte, flag bool) {
	if flag {
		st.indexY = imm
	} else {
		st.indexY = st.memory[loc]
	}
}

func (st *State) lsr() {
	// TODO
}

func (st *State) nop() {
	// TODO
}

func (st *State) ora() {
	// TODO
}

func (st *State) pha() {
	st.stackPush(st.accumulator)
}

func (st *State) php() {
	var value byte = (st.carry << 0) + (st.zero << 1) + (st.interrupt << 2) +
		(st.decimal << 3) + (st.breakFlag << 4) + (1 << 5) + (st.overflow << 6) +
		(st.negative << 7)

	st.stackPush(value)
}

func (st *State) pla() {
	st.accumulator = st.stackPull()
}

func (st *State) plp() {
	value := st.stackPull()

	st.carry = value >> 0
	st.zero = value >> 1
	st.interrupt = value >> 2
	st.decimal = value >> 3
	st.breakFlag = value >> 4
	st.overflow = value >> 6
	st.negative = value >> 7
}

func (st *State) rol() {
	// TODO
}

func (st *State) ror() {
	// TODO
}

func (st *State) rti() {
	// TODO
}

func (st *State) rts() {
	value := uint16(st.stackPull())      // Pull lo byte first
	value += uint16(st.stackPull() >> 8) // Pull hi byte next
	st.ProgramCounter = value + 1        // next opcode is at value + 1
}

func (st *State) sbc() {
	// TODO
}

func (st *State) sec() {
	st.carry = 1
}

func (st *State) sed() {
	st.decimal = 1
}

func (st *State) sei() {
	st.interrupt = 1
}

func (st *State) sta(loc uint16) {
	st.memory[loc] = st.accumulator
}

func (st *State) stx(loc uint16) {
	st.memory[loc] = st.indexX
}

func (st *State) sty(loc uint16) {
	st.memory[loc] = st.indexY
}

func (st *State) tax() {
	st.indexX = st.accumulator
}

func (st *State) tay() {
	st.indexY = st.accumulator
}

func (st *State) tsx() {
	st.indexX = st.stackPointer
}

func (st *State) txa() {
	st.accumulator = st.indexX
}

func (st *State) txs() {
	st.stackPointer = st.indexX
}

func (st *State) tya() {
	st.accumulator = st.indexY
}
