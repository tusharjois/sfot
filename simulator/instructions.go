package simulator

func (st *State) adc(loc uint16, imm byte, flag bool) {
	if flag {
		addition := uint16(accumulator) + uint16(imm)
	} else {
		addition := uint16(accumulator) + loc
	}

	bitSeven := (addition >> 7) & 1

	st.negative = bitSeven

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
		value = byte(loc)
	}

	st.accumulator = st.accumulator & value

	st.negative = (st.accumulator >> 7) & 1

	if st.accumulator == 0 {
		st.zero = 1
	} else {
		st.zero = 0
	}
}

func (st *State) asl(loc uint16, imm byte, flag bool) {
	// TODO
}

func (st *State) bcc(loc uint16) {
	if st.carryFlag == 0 {
		st.ProgramCounter = loc
	}
}

func (st *State) bcs(loc uint16) {
	if st.carryFlag == 1 {
		st.ProgramCounter = loc
	}
}

func (st *State) beq(loc uint16) {
	if st.zero == 1 {
		st.ProgramCounter = loc
	}
}

func (st *State) bit() {}

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

func (st *State) cmp(loc uint16, imm byte, flag bool) {}

func (st *State) cpx() {}

func (st *State) cpy() {}

func (st *State) dec() {}

func (st *State) dex() {}

func (st *State) dey() {}

func (st *State) eor() {}

func (st *State) inc() {}

func (st *State) inx() {}

func (st *State) iny() {}

func (st *State) jmp() {}

func (st *State) jsr() {}

func (st *State) lda() {}

func (st *State) ldx() {}

func (st *State) ldy() {}

func (st *State) lsr() {}

func (st *State) nop() {}

func (st *State) ora() {}

func (st *State) pha() {}

func (st *State) php() {}

func (st *State) pla() {}

func (st *State) plp() {}

func (st *State) rol() {}

func (st *State) ror() {}

func (st *State) rti() {}

func (st *State) rts() {}

func (st *State) sbc() {}

func (st *State) sec() {}

func (st *State) sed() {}

func (st *State) sei() {}

func (st *State) sta() {}

func (st *State) stx() {}

func (st *State) sty() {}

func (st *State) tax() {}

func (st *State) tay() {}

func (st *State) tsx() {}

func (st *State) txa() {}

func (st *State) txs() {}

func (st *State) tya() {}
