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

type LocalAssignNode struct {
	index int
	expr  Node
}

func (n LocalAssignNode) Eval(ctx Ctx) Value {
	ctx.SetLocal(n.index, n.expr.Eval(ctx))
	return Null
}

type VarNode struct {
	key Key
}

func (n VarNode) Eval(ctx Ctx) Value {
	return ctx.Get(n.key)
}

var _ Node = VarNode{}

type LocalVarNode struct {
	index int
}

func (n LocalVarNode) Eval(ctx Ctx) Value {
	return ctx.GetLocal(n.index)
}

var _ Node = LocalVarNode{}

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
	cond, then, else_ Node
}

type EqNode struct {
	left, right Node
}

func (n EqNode) Eval(ctx Ctx) Value {
	return ToBool(n.left.Eval(ctx).Compare(n.right.Eval(ctx)).(Int) == Int(0))
}

var _ Node = EqNode{}

type GtNode struct {
	left, right Node
}

func (n GtNode) Eval(ctx Ctx) Value {
	return ToBool(n.left.Eval(ctx).Compare(n.right.Eval(ctx)).(Int) == Int(1))
}

func (n IfNode) Eval(ctx Ctx) Value {
	if n.cond.Eval(ctx).True() {
		return n.then.Eval(ctx)
	} else {
		return n.else_.Eval(ctx)
	}
}

type ForNode struct {
	index int
	iter  Value
	code  Node
}

func (n ForNode) Eval(ctx Ctx) Value {
	iter := n.iter.Iter()
	for iter.Next() {
		ctx.SetLocal(n.index, iter.Value())
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
	subCtx := fromCtx(ctx, []Value{})
	for i, k := range n.argNames {
		subCtx.Set(k, n.argValues[i].Eval(ctx))
	}
	return n.callable.Eval(subCtx)
}
