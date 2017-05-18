package govm

type Key string

type Ctx interface {
	Get(key Key) Value
	Set(key Key, value Value)
}

type Node interface {
	Eval(ctx Ctx) Value
}

type Value interface {
	Node

	True() bool

	// arithmetic
	Add(other Value) Value
	Sub(other Value) Value
	// TODO
	// Mul(other Value) Value
	// Div(other Value) Value

	// comparizon
	Compare(other Value) Value

	GetItem(index Value) Value
	SetItem(index, value Value)

	// iterable
	Next() (Value, bool)
}
