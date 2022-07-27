package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) GetInScope(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	_, ok := e.GetInScope(name)
	if !ok {
		if e.outer != nil {
			return e.outer.Set(name, val)
		} else {
			return nil
		}
	} else {
		return e.SetInScope(name, val)
	}
}

func (e *Environment) SetInScope(name string, val Object) Object {
	e.store[name] = val
	return val
}
