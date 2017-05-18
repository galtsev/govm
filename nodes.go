package govm

type AssignNode struct {
	key  Key
	expr Node
}

func (n AssignNode) Eval(ctx Ctx) Value {
	ctx.Set(n.key, n.expr.Eval(ctx))
	return Null
}

var _ Node = AssignNode{}

type AddNode struct {
	left, right Node
}

func (n AddNode) Eval(ctx Ctx) Value {
	return n.left.Eval(ctx).Add(n.right.Eval(ctx))
}

type SubNode struct {
	left, right Node
}

func (n SubNode) Eval(ctx Ctx) Value {
	return n.left.Eval(ctx).Sub(n.right.Eval(ctx))
}

type IfNode struct {
	cond, then, _else Node
}

func (n IfNode) Eval(ctx Ctx) Value {
	if n.cond.Eval(ctx).True() {
		return n.then.Eval(ctx)
	} else {
		return n._else.Eval(ctx)
	}
}

type ForNode struct {
	key  Key
	iter Value
	code Node
}

func (n ForNode) Eval(ctx Ctx) Value {
	for v, ok := n.iter.Next(); ok; v, ok = n.iter.Next() {
		ctx.Set(n.key, v)
		n.code.Eval(ctx)
	}
	return Null
}

type CallNode struct {
	callable  Node
	argNames  []Key
	argValues []Node
}

func (n CallNode) Eval(ctx Ctx) Value {
	subCtx := fromCtx(ctx)
	for i, k := range n.argNames {
		subCtx.Set(k, n.argValues[i].Eval(ctx))
	}
	return n.callable.Eval(subCtx)
}
