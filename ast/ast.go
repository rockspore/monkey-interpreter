package ast

import (
	"bytes"

	"github.com/rockspore/monkey-interpreter/token"
)

// Node : the most basic interface for ast node
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement : statement node
type Statement interface {
	Node
	statementNode()
}

// Expression : expression node
type Expression interface {
	Node
	expressionNode()
}

// Program : root of an ast for a program containing statements
type Program struct {
	Statements []Statement
}

// TokenLiteral : returns the first statement's token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String : returns each statement's string form of the program
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement : lets tatement implementing the statement interface
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// guide to Go compiler that LetStatement implements Statement interface
func (ls *LetStatement) statementNode() {}

// TokenLiteral : returns the let token literal
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String : returns the string form of the let statement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement : return statement implementing the statement interface
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

// guide to Go compiler that ReturnStatement implements Statement interface
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral : returns the return token literal
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String : returns the string form of the return statement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement : a statement containing solely of one expression
type ExpressionStatement struct {
	Token      token.Token // the first token fo the expression
	Expression Expression
}

// guide to Go compiler that ExpressionStatement implements Statement interface
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral : returns the return token literal
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String : returns the string form of the expression statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier : implements the Expression interface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// guide to Go compiler that Identifier implements Expression interface
func (i *Identifier) expressionNode() {}

// TokenLiteral : returns the identifier's token literal
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String : returns the string form of the identifier
func (i *Identifier) String() string {
	return i.Value
}
