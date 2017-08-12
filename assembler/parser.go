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

		if lbl, prs := ctx.labels[lbl.Label]; prs {
			// Absolute addressing
			return NewInstrNodeFromLabel(instrName, lbl, "abs")
		}

		ln = new(labelNode)
		ln.content = labelName

		// Absolute addressing
		return NewInstrNodeFromLabel(instrName, ln, "abs")

	case "Instruction":
		// Implicit addressing
		return NewInstrNodeFromAddr(instrName, 0x0, "imp")
	}

}

func NewContext(tz *Tokenizer) *context {
	return &context{tz: tz, current: tz.Next(), last: nil, labels: make(map[string]*labelNode)}
}

func Parse(tz *Tokenizer) ([]node, error) {
	var nodeList []node

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
				return nodeList, err
			}
		} else {
			if i, err := ctx.handleInstruction(); err == nil {
				nodeList = append(nodeList, i)
			} else {
				return nodeList, err
			}
		}
	}

	return nodeList, nil
}
