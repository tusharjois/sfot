package main

import (
	"fmt"
	"github.com/tusharjois/sfot/assembler"
)

func main() {

	str := "\tPHA\n\tLDA $02,X\nlbl:\n\tASL A ; Here's a comment\n\t; Here's another one\n\tSTA ($00, X)\nJMP lbl\n"
	/*str = `
	  LDX #$00
	  LDY #$00
	firstloop:
	  TXA
	  STA $0200,Y
	  PHA
	  INX
	  INY
	  CPY #$10
	  BNE firstloop ;loop until Y is $10
	secondloop:
	  PLA
	  STA $0200,Y
	  INY
	  CPY #$20      ;loop until Y is $20
	  BNE secondloop
	  `*/
	fmt.Println(str)
	tz, err := assembler.NewTokenizer(&str)

	if err != nil {
		fmt.Println(err)
	} else {
		p, err := assembler.Parse(tz)

		if err != nil {
			fmt.Println(err)
		} else {
			assembler.Assemble(p)
		}
	}

}
