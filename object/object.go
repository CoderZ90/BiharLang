package object

import "fmt"

type ObjectType string

const (
	INTEGER      ObjectType = "INTEGER"
	STRING       ObjectType = "STRING"
	BOOLEAN      ObjectType = "BOOLEAN"
	NULL         ObjectType = "NULL"
	RETURN_VALUE ObjectType = "RETURN_VALUE"
	ERROR        ObjectType = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

//  Integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

//  String
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING }
func (s *String) Inspect() string  { return s.Value }

//  Boolean
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

//  Null (kuchnahi)
type Null struct{}

func (n *Null) Type() ObjectType { return NULL }
func (n *Null) Inspect() string  { return "kuchnahi" }

// Error
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR }
func (e *Error) Inspect() string  { return "💀 ERROR: " + e.Message }
