package assembler

import (
	"errors"
	"fmt"
)

type context struct {
	tz      *Tokenizer
	current *Token
	last    *Token
	labels  map[string]*labelNode
}

func (ctx *context) match(toMatch ...string) (*Token, error) {
	for _, element := range toMatch {
		if element == ctx.current.Kind {
			ctx.next()
			return ctx.last, nil
		}
	}

	msg := fmt.Sprintf("invalid token - expected %v, got %v", toMatch[0], ctx.current)
	err := errors.New(msg)
	return nil, err
}

func (ctx *context) next() {
	ctx.last = ctx.current
	ctx.current = ctx.tz.Next()
}

func (ctx *context) handleLabel() (*labelNode, error) {
	labelName := ctx.last.Label

	if _, err := ctx.match(":"); err != nil {
		return nil, err
	}

	var ln *labelNode

	if curr, prs := ctx.labels[labelName]; prs {
		curr.content = labelName
		ln = curr
	} else {
		ln = new(labelNode)
		ln.content = labelName
	}

	return ln, nil
}

func (ctx *context) handleInstruction() (*instrNode, error) {
	instrName := ctx.last.Label
	switch ctx.current.Kind {
	case "A":
		ctx.match("A")
		// Accumulator addressing
		return NewInstrNodeFromAddr(instrName, 0x0, "acc")
	case "#":
		if _, err := ctx.match("#"); err != nil {
			return nil, err
		}
		num, err := ctx.match("Number")
		if err != nil {
			return nil, err
		}
		// Immediate addressing
		return NewInstrNodeFromAddr(instrName, num.Number, "imm")
	case "Label":
		lbl, err := ctx.match("Label")
		if err != nil {
			return nil, err
		}

		// TODO: relative branching
		if lbl, prs := ctx.labels[lbl.Label]; prs {
			// Absolute addressing
			return NewInstrNodeFromLabel(instrName, lbl, "abs")
		}

		ln := new(labelNode)
		ln.content = lbl.Label
		ctx.labels[lbl.Label] = ln

		// Absolute addressing
		return NewInstrNodeFromLabel(instrName, ln, "abs")
	case "Number":
		num, err := ctx.match("Number")
		if err != nil {
			return nil, err
		}

		if ctx.current.Kind == "," {
			_, err := ctx.match(",")
			if err != nil {
				return nil, err
			}
			reg, err := ctx.match("X", "Y")
			if err != nil {
				return nil, err
			}

			if num.Label == "uint8" {
				num.Label = "zp"
			} else {
				num.Label = "ab"
			}

			// Absolute,R OR Zero Page,R addressing
			if reg.Kind == "X" {
				return NewInstrNodeFromAddr(instrName, num.Number, num.Label+"x")
			}
			return NewInstrNodeFromAddr(instrName, num.Number, num.Label+"y")
		}

		if num.Label == "uint8" {
			// Zero Page addressing
			return NewInstrNodeFromAddr(instrName, num.Number, "zp0")
		}
		// Absolute addressing
		return NewInstrNodeFromAddr(instrName, num.Number, "abs")
		// TODO: indirect addressing

	}
	// Implicit addressing
	return NewInstrNodeFromAddr(instrName, 0x0, "imp")

}

func NewContext(tz *Tokenizer) *context {
	return &context{tz: tz, current: tz.Next(), last: nil, labels: make(map[string]*labelNode)}
}

func Parse(tz *Tokenizer) ([]Node, error) {
	var nodeList []Node

	ctx := NewContext(tz)
	for ctx.current != nil {
		t, err := ctx.match("Instruction", "Label")
		if err != nil {
			return nodeList, err
		}
		if t.Kind == "Label" {
			if l, err := ctx.handleLabel(); err == nil {
				nodeList = append(nodeList, l)
			} else {
				return nodeList,
					errors.New(fmt.Sprintf("%v: ", t) + err.Error())
			}
		} else {
			if i, err := ctx.handleInstruction(); err == nil {
				nodeList = append(nodeList, i)
			} else {
				return nodeList,
					errors.New(fmt.Sprintf("%v: ", t) + err.Error())
			}
		}
	}

	return nodeList, nil
}
