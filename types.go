package govm

import (
	"errors"
)

// True() bool

// // arithmetic
// Add(other Value) Value
// Sub(other Value) Value
// Mul(other Value) Value
// Div(other Value) Value

// // comparizon
// Compare(other Value) int

// GetItem(index Value) Value
// SetItem(index, value Value)

// Call(args []Value) Value

// // iterable
// Next() (Value, bool)

var (
	NotImplemented  = errors.New("Not implemented")
	ConversionError = errors.New("Conversion error")
)

type Stub struct{}

func (s Stub) Eval(ctx Ctx) Value               { panic(NotImplemented) }
func (s Stub) True() bool                       { panic(NotImplemented) }
func (s Stub) Add(other Value) Value            { panic(NotImplemented) }
func (s Stub) Sub(other Value) Value            { panic(NotImplemented) }
func (s Stub) Mul(other Value) Value            { panic(NotImplemented) }
func (s Stub) Div(other Value) Value            { panic(NotImplemented) }
func (s Stub) Compare(other Value) Value        { panic(NotImplemented) }
func (s Stub) GetItem(index Value) Value        { panic(NotImplemented) }
func (s Stub) SetItem(index, value Value)       { panic(NotImplemented) }
func (s Stub) Call(args []Value, ctx Ctx) Value { panic(NotImplemented) }
func (s Stub) Next() (Value, bool)              { panic(NotImplemented) }

type nullType struct {
	Stub
}

var Null Value = nullType{}

type Int int

func ToInt(v Value) Int {
	if v, ok := v.(Int); ok {
		return v
	}
	panic(ConversionError)
}

func (i Int) Eval(ctx Ctx) Value    { return i }
func (i Int) True() bool            { return i != 0 }
func (i Int) Add(other Value) Value { return i + ToInt(other) }
func (i Int) Sub(other Value) Value { return i - ToInt(other) }
func (i Int) Mul(other Value) Value { return i * ToInt(other) }
func (i Int) Div(other Value) Value { return i / ToInt(other) }
func (i Int) Compare(other Value) Value {
	d := i - ToInt(other)
	if d > 0 {
		return Int(1)
	} else if d < 0 {
		return Int(-1)
	} else {
		return Int(0)
	}
}
func (i Int) GetItem(index Value) Value        { panic(NotImplemented) }
func (i Int) SetItem(index, value Value)       { panic(NotImplemented) }
func (i Int) Call(args []Value, ctx Ctx) Value { panic(NotImplemented) }
func (i Int) Next() (Value, bool)              { panic(NotImplemented) }

var _ Value = Int(0)
