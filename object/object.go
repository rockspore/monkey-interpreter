package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rockspore/monkey-interpreter/ast"
)

// ObjectType : object type
type ObjectType string

const (
	IntegerOBJ     = "INTEGER"
	BooleanOBJ     = "BOOLEAN"
	NullOBJ        = "NULL"
	ReturnValueOBJ = "RETURN_VALUE"
	ErrorOBJ       = "ERROR"
	FunctionOBJ    = "FUNCTION"
)

// Object : object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer : integer object
type Integer struct {
	Value int64
}

// Inspect : return integer value
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type : return integer type
func (i *Integer) Type() ObjectType {
	return IntegerOBJ
}

// Boolean : boolean object
type Boolean struct {
	Value bool
}

// Inspect : return boolean value
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type : return boolean type
func (b *Boolean) Type() ObjectType {
	return BooleanOBJ
}

// Null : null object
type Null struct{}

// Inspect : return null
func (n *Null) Inspect() string {
	return "null"
}

// Type : return null type
func (n *Null) Type() ObjectType {
	return NullOBJ
}

// ReturnValue : return value object
type ReturnValue struct {
	Value Object
}

// Inspect : return the returned object value
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Type : return return-value type
func (rv *ReturnValue) Type() ObjectType {
	return ReturnValueOBJ
}

// Error : error object
type Error struct {
	Message string
}

// Inspect : return error message
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Type : return error object type
func (e *Error) Type() ObjectType {
	return ErrorOBJ
}

// Function : function object
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Inspect : return function object
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// Type : return function type
func (f *Function) Type() ObjectType {
	return FunctionOBJ
}
