package body

type Body []byte

func NewBody() Body {
	return make(Body, 0)
}

func (b *Body) Parse(lines []byte) []byte {
	*b = lines
	return *b
}
