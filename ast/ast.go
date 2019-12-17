package ast

import "github.com/rockspore/monkey-interpreter/token"

// Node : tbd
type Node interface {
	TokenLiteral() string
}

// Statement : tbd
type Statement interface {
	Node
	statementNode()
}

// Expression : tbd
type Expression interface {
	Node
	expressionNode()
}

// Program : tbd
type Program struct {
	Statements []Statement
}

// TokenLiteral : tbd
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement : implements the Statement interface
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// guide to Go compiler that LetStatement implements Statement interface
func (ls *LetStatement) statementNode() {}

// TokenLiteral : tbd
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier : implements the Expression interface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// guide to Go compiler that Identifier implements Expression interface
func (i *Identifier) expressionNode() {}

// TokenLiteral : tbd
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
