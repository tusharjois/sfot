package assembler

type context struct {
	tz      *Tokenizer
	current *Token
	last    *Token
}

func (ctx *context) match(toMatch ...string) *Token {
	for _, element := range toMatch {
		if element == p.current.Kind {
			pc.next()
			return
		}
	}
}

func (ctx *context) next() {
	last = current // For semantic analysis
	current = tz.Next()
}

func NewContext(tz *Tokenizer) *ParserContext {
	return nil
}
