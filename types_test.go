package govm

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const N int = 10

func toInt(value Value) int {
	return int(value.(Int))
}

func TestInt(t *testing.T) {
	ctx := BaseCtx()
	for i := 0; i < N; i++ {
		a := int(rand.Int63())
		b := int(rand.Int63())
		va := Int(a)
		vb := Int(b)
		assert.Equal(t, a, toInt(va.Eval(ctx)))
		assert.Equal(t, a+b, toInt(va.Add(vb)))
		assert.Equal(t, a-b, toInt(va.Sub(vb)))
		var expectCompare int
		if a > b {
			expectCompare = 1
		} else if a < b {
			expectCompare = -1
		} else {
			expectCompare = 0
		}
		assert.Equal(t, expectCompare, toInt(va.Compare(vb)))
	}
}

func TestRange(t *testing.T) {
	n := rand.Intn(N)
	rng := Range(Int(n))
	cnt := 0
	for rng.Next() {
		cnt++
	}
	assert.Equal(t, n, cnt)
}
