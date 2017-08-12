package main

import (
	"fmt"
	"github.com/tusharjois/sfot/tokenizer"
)

func main() {

	str := "\tPHA\n\tLDA #$02\nlbl:\n\tCLC\n\tSTA ($00, X)\n"
	fmt.Println(str)
	tz, err := tokenizer.NewTokenizer(&str)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tz)
	}
}
