package simulator

func (st *State) Run() {
	isRunning := st.Step()

	for isRunning {
		isRunning = st.Step()
	}
}

func (st *State) Step() bool {
	opcode := st.memory[st.ProgramCounter]
	st.ProgramCounter++

	info := strings.Split(opcodeMatrix[opcode], "-")
	// info[0] is the name of the opcode, info[1] is the addressing mode

	var location uint16
	var immediate byte
	var immediateFlag bool = false

	// TODO: Panic Checks
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
		location += uint32(st.memory[st.ProgramCounter]) << 8
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
		location += uint16(st.memory[indirectLoc]) << 8
	case "inx":
		indirectLoc := uint16(st.memory[st.ProgramCounter]) + uint16(st.indexX)
		st.ProgramCounter++
		if lookupLoc > 0xff {
			lookupLoc &= 0xfe
		}
		location = uint16(st.memory[indirectLoc])
		indirectLoc++
		location += uint16(st.memory[indirectLoc]) << 8
	case "iny":
		indirectLoc := st.ProgramCounter
		st.ProgramCounter++
		location = uint16(st.memory[indirectLoc])
		indirectLoc++
		location += uint16(st.memory[indirectLoc]) << 8
		tempLoc := uint32(st.indexY) + uint32(location)
		if tempLoc > 0xffff {
			tempLoc -= 0xffff
		}
		location = uint16(tempLoc)
	case "rel":
		// TODO
	case "imp":
	}

	st.handleInstruction(info[0], location, immediate, isImm)
}

func (st *State) handleInstruction(instr string, location uint16, immediate byte, flag bool) {

	switch instr {
	case "ADC":
		adc(location, immediate, flag)
	case "AND":
		and(location, immediate, flag)
	case "ASL":
		asl()
	case "BCC":
		bcc()
	case "BCS":
		bcs()
	case "BEQ":
		beq()
	case "BIT":
		bit()
	case "BMI":
		bmi()
	case "BNE":
		bne()
	case "BPL":
		bpl()
	case "BRK":
		brk()
	case "BVC":
		bvc()
	case "BVS":
		bvs()
	case "CLC":
		clc()
	case "CLD":
		cld()
	case "CLI":
		cli()
	case "CLV":
		clv()
	case "CMP":
		cmp()
	case "CPX":
		cpx()
	case "CPY":
		cpy()
	case "DEC":
		dec()
	case "DEX":
		dex()
	case "DEY":
		dey()
	case "EOR":
		eor()
	case "INC":
		inc()
	case "INX":
		inx()
	case "INY":
		iny()
	case "JMP":
		jmp()
	case "JSR":
		jsr()
	case "LDA":
		lda()
	case "LDX":
		ldx()
	case "LDY":
		ldy()
	case "LSR":
		lsr()
	case "NOP":
		nop()
	case "ORA":
		ora()
	case "PHA":
		pha()
	case "PHP":
		php()
	case "PLA":
		pla()
	case "PLP":
		plp()
	case "ROL":
		rol()
	case "ROR":
		ror()
	case "RTI":
		rti()
	case "RTS":
		rts()
	case "SBC":
		sbc()
	case "SEC":
		sec()
	case "SED":
		sed()
	case "SEI":
		sei()
	case "STA":
		sta()
	case "STX":
		stx()
	case "STY":
		sty()
	case "TAX":
		tax()
	case "TAY":
		tay()
	case "TSX":
		tsx()
	case "TXA":
		txa()
	case "TXS":
		txs()
	case "TYA":
		tya()
	}
}
