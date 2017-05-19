package govm

type Context struct {
	parent Ctx
	values map[Key]Value
	locals []Value
}

func (ctx *Context) Get(key Key) Value {
	if v, ok := ctx.values[key]; ok {
		return v
	}
	if ctx.parent != nil {
		return ctx.parent.Get(key)
	}
	return Null
}

func (ctx *Context) Set(key Key, value Value) {
	ctx.values[key] = value
}

func (ctx *Context) GetLocal(index int) Value {
	return ctx.locals[index]
}

func (ctx *Context) SetLocal(index int, value Value) {
	ctx.locals[index] = value
}

func BaseCtx() Ctx {
	return &Context{
		values: make(map[Key]Value),
	}
}

func fromCtx(parent Ctx, locals []Value) Ctx {
	return &Context{
		parent: parent,
		values: make(map[Key]Value),
		locals: locals,
	}
}
