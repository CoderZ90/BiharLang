package evaluator

import "bihari/object"

type Environment struct {
	store map[string]object.Object
}

// Init new environment, store={}
func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]object.Object),
	}
}

// putting something in the environment
func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}

//getting a value back
func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
