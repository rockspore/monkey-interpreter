package object

import (
	"fmt"
)

// ObjectType : object type
type ObjectType string

const (
	IntegerOBJ = "INTEGER"
	BooleanOBJ = "BOOLEAN"
	NullOBJ    = "NULL"
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
