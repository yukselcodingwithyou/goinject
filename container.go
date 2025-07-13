package goinject

import (
	"errors"
	"fmt"
	"reflect"
)

// Container holds provided dependencies and can inject them into structs.
type entry struct {
	name  string
	value reflect.Value
}

// Container holds provided dependencies and can inject them into structs.
type Container struct {
	values []entry
}

// New creates a new dependency container.
func New() *Container {
	return &Container{values: make([]entry, 0)}
}

// Provide registers a dependency instance with the container.
func (c *Container) Provide(v interface{}) {
	c.ProvideNamed("", v)
}

// ProvideNamed registers a dependency with a qualifier name. The empty name is
// used for default injection when no qualifier is specified.
func (c *Container) ProvideNamed(name string, v interface{}) {
	c.values = append(c.values, entry{name: name, value: reflect.ValueOf(v)})
}

// get finds a value assignable to t.
func (c *Container) get(t reflect.Type, name string) (reflect.Value, bool) {
	for _, e := range c.values {
		if name != "" && e.name != name {
			continue
		}
		v := e.value
		if v.Type().AssignableTo(t) {
			return v, true
		}
		if v.CanAddr() && v.Addr().Type().AssignableTo(t) {
			return v.Addr(), true
		}
	}
	return reflect.Value{}, false
}

// Fill populates fields tagged with `autowire` using dependencies from the container.
// The tag value can optionally specify a qualifier name registered with ProvideNamed.
func (c *Container) Fill(target interface{}) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New("target must be pointer to struct")
	}

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		structField := rt.Field(i)

		tagVal, ok := structField.Tag.Lookup("autowire")
		if !ok {
			continue
		}

		if !field.CanSet() {
			return fmt.Errorf("cannot set field %s", structField.Name)
		}

		depVal, found := c.get(structField.Type, tagVal)
		if !found {
			return fmt.Errorf("no dependency for field %s of type %s", structField.Name, structField.Type)
		}

		field.Set(depVal)
	}
	return nil
}
