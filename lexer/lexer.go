package lexer

// Lexer : definition of the lexer struct
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New : generates a new lexer based on the input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
