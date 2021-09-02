package object

// NewEnvironment : create a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment : create a new environment extending the outer env
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Environment : the envrionment object
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get : return the object with the specified identifier from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set : set and return the objected with specified identifier
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
