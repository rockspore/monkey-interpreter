package parser

import (
	"github.com/rockspore/monkey-interpreter/ast"
	"github.com/rockspore/monkey-interpreter/lexer"
	"github.com/rockspore/monkey-interpreter/token"
)

// Parser : tbd
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New : tbd
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so both curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.nextToken()
}

// ParseProgram : tbd
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statements{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
