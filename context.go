package govm

type Context struct {
	parent Ctx
	values map[Key]Value
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

func fromCtx(parent Ctx) Ctx {
	return &Context{
		parent: parent,
		values: make(map[Key]Value),
	}
}
