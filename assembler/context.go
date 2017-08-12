package assembler

type context struct {
	tz      *Tokenizer
	current *Token
	last    *Token
	labels  map[string]*label
}

func (ctx *context) match(toMatch ...string) (*Token, error) {
	for _, element := range toMatch {
		if element == ctx.current.Kind {
			ctx.next()
			return ctx.last, nil
		}
	}

	msg := fmt.Sprintf("invalid token - expected %v, got %v", toMatch[0], current)
	err := errors.New(msg)
	return nil, err
}

func (ctx *context) next() {
	ctx.last = ctx.current
	ctx.current = ctx.tz.Next()
}

func NewContext(tz *Tokenizer) *context {
	return &context{tz: tz, current: tz.Next(), last: nil, labels: make(map[string]*label)}
}
