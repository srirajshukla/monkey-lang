package object

import (
	"bytes"
	"fmt"
	"gks/monkey_intp/ast"
	"strings"
)

type ObjectType string

const (
	INTEJER_OBJ      = "INTEGER"
	STRING_OBJ       = "STRING"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETRUN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
)

type Object interface {
	Type() ObjectType
	Insepct() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Insepct() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEJER_OBJ
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Insepct() string {
	return s.Value
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Insepct() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type Null struct{}

func (n *Null) Insepct() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Insepct() string {
	return rv.Value.Insepct()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Insepct() string {
	return "ERROR: " + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) Insepct() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn (")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

func (b *Builtin) Insepct() string {
	return "builtin function"
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (a *Array) Insepct() string {
	var out bytes.Buffer

	elements := []string{}

	for _, e := range a.Elements {
		elements = append(elements, e.Insepct())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}
