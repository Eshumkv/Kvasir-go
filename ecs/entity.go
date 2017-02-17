package ecs

import (
	"reflect"
	"sync"
	"sync/atomic"
)

var (
	counterLock sync.Mutex
	idInc       uint64
)

// Entity defines the entity in the entity component system.
type Entity struct {
	id         uint64
	x, y       float64
	w, h       uint16
	components []Component
}

// NewEntity creates a new entity with a unique identifier.
func NewEntity(x, y float64, w, h uint16) *Entity {
	return &Entity{
		id:         atomic.AddUint64(&idInc, 1),
		x:          x,
		y:          y,
		w:          w,
		h:          h,
		components: make([]Component, 0),
	}
}

// ID returns the unique id of the entity.
func (e Entity) ID() uint64 {
	return e.id
}

// Add a component to the entity.
func (e *Entity) Add(comp Component) *Entity {
	e.components = append(e.components, comp)
	return e
}

// Remove a component from the entity.
func (e *Entity) Remove(comp string) {
	index := -1
	for i, c := range e.components {
		if reflect.TypeOf(c).Name() == comp {
			index = i
			break
		}
	}

	if index != -1 {
		copy(e.components[index:], e.components[index+1:])
		e.components[len(e.components)-1] = nil
		e.components = e.components[:len(e.components)-1]
	}
}

// Get the component that you want, returns (component, ok)
func (e Entity) Get(comp string) (Component, bool) {
	for _, c := range e.components {
		if reflect.TypeOf(c).Name() == comp {
			return c, true
		}
	}
	return nil, false
}

// Has the entity a specific component?
func (e Entity) Has(comp string) bool {
	for _, c := range e.components {
		if reflect.TypeOf(c).Name() == comp {
			return true
		}
	}
	return false
}
