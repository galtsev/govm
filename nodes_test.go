package govm

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestAssignNode(t *testing.T) {
	ctx := BaseCtx()
	k := Key("v")
	v := int(rand.Int63())
	code := AssignNode{
		key:  k,
		expr: Int(v),
	}
	code.Eval(ctx)
	assert.Equal(t, v, toInt(ctx.Get(k)))
}

func TestVarNode(t *testing.T) {
	ctx := BaseCtx()
	k := Key("v")
	v := int(rand.Int63())
	ctx.Set(k, Int(v))
	code := VarNode{
		key: k,
	}
	assert.Equal(t, v, toInt(code.Eval(ctx)))
}

func TestIfNode(t *testing.T) {
	ctx := BaseCtx()
	then := Int(rand.Int63())
	else_ := Int(rand.Int63())
	code := IfNode{
		cond:  Int(0), // false
		then:  then,
		else_: else_,
	}
	assert.Equal(t, else_, code.Eval(ctx))
	code.cond = Int(1) // true
	assert.Equal(t, then, code.Eval(ctx))
}

func testBinary(t *testing.T, f func(a, b Value) (expected Value, code Node)) {
	ctx := BaseCtx()
	for i := 0; i < N; i++ {
		a := Int(rand.Int63())
		b := Int(rand.Int63())
		expected, code := f(a, b)
		assert.Equal(t, expected, code.Eval(ctx))
	}
}

func TestAddNode(t *testing.T) {
	testBinary(t, func(a, b Value) (Value, Node) {
		return a.Add(b),
			AddNode{
				left:  a,
				right: b,
			}
	})
}

func TestEqNode(t *testing.T) {
	testBinary(t, func(a, b Value) (Value, Node) {
		return ToBool(a == b),
			EqNode{
				left:  a,
				right: b,
			}
	})
}

// func TestEqNode(t *testing.T) {
// 	ctx := BaseCtx()
// 	for i := 0; i < N; i++ {
// 		a := Int(rand.Int63())
// 		b := Int(rand.Int63())
// 		code := EqNode{
// 			left:  a,
// 			right: b,
// 		}
// 		assert.Equal(t, a == b, code.Eval(ctx).True())
// 	}
// }

func BenchmarkFor(b *testing.B) {
	base := BaseCtx()
	locals := []Value{Int(0), Int(0)}
	ctx := fromCtx(base, locals)
	// variables
	s := 0
	i := 1
	code := ForNode{
		index: i,
		iter:  Range(Int(b.N)),
		code: LocalAssignNode{
			index: s,
			expr: AddNode{
				left: LocalVarNode{
					index: s,
				},
				right: Int(1),
			},
		},
	}
	b.ResetTimer()
	for z := 0; z < b.N; z++ {
		code.Eval(ctx)
	}
	v := ctx.GetLocal(s).(Int)
	assert.Equal(b, b.N, int(v))
}
