package main

import (
	"fmt"
	"github.com/tusharjois/sfot/assembler"
)

func main() {

	str := "\tPHA\n\tLDA #$02\nlbl:\n\tASL A ; Here's a comment\n\t; Here's another one\n\tSTA ($00, X)\n"
	fmt.Println(str)
	tz, err := assembler.NewTokenizer(&str)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tz)
	}

	p, err := assembler.Parse(tz)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}
}
