package lexer

type Lexer struct {
	input string
	// ch is the current character being scanned.
	ch byte
	// position is the current position in input (i.e. index of ch in input).
	position int
	// readPosition is the current reading position in input (i.e. index of character following ch in input).
	readPosition int
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
