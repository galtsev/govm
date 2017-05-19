package govm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
