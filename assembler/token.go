package assembler

import (
	"fmt"
)

type Token struct {
	Kind     string
	Number   uint16
	Label    string
	StartPos int
	EndPos   int
}

func (t *Token) String() string {
	if t.Kind == "Label" {
		return fmt.Sprintf("label<%v>@(%v,%v)", t.Label, t.StartPos, t.EndPos)
	} else if t.Kind == "Number" {
		return fmt.Sprintf("number<%v%#x>@(%v,%v)", t.Label, t.Number, t.StartPos, t.EndPos)
	} else if t.Kind == "Instruction" {
		return fmt.Sprintf("%v@(%v,%v)", t.Label, t.StartPos, t.EndPos)
	}

	return fmt.Sprintf("%v@(%v,%v)", t.Kind, t.StartPos, t.EndPos)
}

func newToken(kind string, number uint16, label string, start int, end int) *Token {
	return &Token{Kind: kind, Number: number, Label: label, StartPos: start,
		EndPos: end}
}
