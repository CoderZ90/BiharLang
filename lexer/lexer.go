package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 here is EOF: end of file
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
