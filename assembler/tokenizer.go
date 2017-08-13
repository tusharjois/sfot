package assembler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Tokenizer struct {
	Tokens   []*Token
	position int
}

func (tz *Tokenizer) String() string {
	output := ""

	for _, t := range tz.Tokens {
		output += fmt.Sprintf("%v\n", t)
	}

	return output
}

func (tz *Tokenizer) Next() *Token {
	var t *Token
	if tz.position < len(tz.Tokens) {
		t = tz.Tokens[tz.position]
		tz.position++
	}
	return t
}

func (tz *Tokenizer) add(t *Token) {
	tz.Tokens = append(tz.Tokens, t)
}

func read(input, c *string, index int) {
	*c = string((*input)[index])
}

func isCharacter(c *string) bool {
	return *c == "#" || *c == "," || *c == "(" || *c == ")" || *c == ":"
}

func isKeyword(c *string) bool {
	keywords := []string{
		"ADC", "AND", "ASL", "BCC", "BCS", "BEQ", "BIT", "BMI",
		"BNE", "BPL", "BRK", "BVC", "BVS", "CLC", "CLD", "CLI",
		"CLV", "CMP", "CPX", "CPY", "DEC", "DEX", "DEY", "EOR",
		"INC", "INX", "INY", "JMP", "JSR", "LDA", "LDX", "LDY",
		"LSR", "NOP", "ORA", "PHA", "PHP", "PLA", "PLP", "ROL",
		"ROR", "RTI", "RTS", "SBC", "SEC", "SED", "SEI", "STA",
		"STX", "STY", "TAX", "TAY", "TSX", "TXA", "TXS", "TYA",
		"X", "Y", "A"}

	for _, word := range keywords {
		if strings.ToUpper(*c) == word {
			return true
		}
	}
	return false
}

func isHex(c *string) bool {
	keywords := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a",
		"b", "c", "d", "e", "f"}

	for _, word := range keywords {
		if word == strings.ToLower(*c) {
			return true
		}
	}
	return false
}

func NewTokenizer(input *string) (*Tokenizer, error) {
	tz := new(Tokenizer)
	l := len(*input)

	index := 0
	var start, end int

	for index < l {
		current := string((*input)[index])
		c := &current

		if *c == "\n" || *c == " " || *c == "\t" {
			index++
		} else if *c == ";" {
			// Comment - ignore rest of line
			for *c != "\n" {
				index++
				if index >= l {
					break
				}
				read(input, c, index)
			}
		} else if isCharacter(c) {
			t := newToken(*c, 0x0, "", index, index)
			tz.add(t)
			index++
		} else if *c == "$" {
			index++
			if index >= l {
				break
			}
			read(input, c, index)
			start = index
			end = index
			var integer []byte
			for isHex(c) {
				end = index
				integer = append(integer, (strings.ToLower(*c))[0])
				index++
				if index >= l {
					break
				}
				read(input, c, index) // advance the automaton
			}
			number, err := strconv.ParseUint(string(integer), 16, 16)
			if err != nil {
				msg := fmt.Sprintf("number at (%v, %v) is not uint16", start, end)
				return nil, errors.New(msg)
			}
			var extra string
			if len(integer) == 2 {
				extra = "uint8:"
			}
			t := newToken("Number", uint16(number), extra, start, end)
			tz.add(t)
		} else if unicode.IsDigit(rune((*c)[0])) {
			start = index
			end = index
			var integer []byte
			for unicode.IsDigit(rune((*c)[0])) {
				end = index
				integer = append(integer, (strings.ToLower(*c))[0])
				index++
				if index >= l {
					break
				}
				read(input, c, index) // advance the automaton
			}
			number, err := strconv.ParseUint(string(integer), 10, 16)
			if err != nil {
				msg := fmt.Sprintf("number at (%v, %v) is not uint16", start, end)
				return nil, errors.New(msg)
			}
			t := newToken("Number", uint16(number), "", start, end)
			tz.add(t)
		} else if unicode.IsLetter(rune((*c)[0])) {
			start = index
			end = index
			var identifier []byte
			for unicode.IsLetter(rune((*c)[0])) || unicode.IsDigit(rune((*c)[0])) {
				end = index
				identifier = append(identifier, (*c)[0])
				index++
				if index >= l {
					break
				}
				read(input, c, index) // advance the automaton
			}
			strRepr := string(identifier)
			if isKeyword(&strRepr) {
				title := "Instruction"
				if strRepr == "X" || strRepr == "A" || strRepr == "Y" {
					title = strRepr
				}
				t := newToken(title, 0, strings.ToUpper(strRepr), start, end)
				tz.add(t)
			} else {
				t := newToken("Label", 0, strRepr, start, end)
				tz.add(t)
			}
		} else {
			msg := fmt.Sprintf("unrecognized character %v at position %v", *c, index)
			return nil, errors.New(msg)
		}

	}
	t := newToken("eof", 0, "", index, index)
	tz.add(t)
	return tz, nil
}
