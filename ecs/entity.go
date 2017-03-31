package ecs

import (
	"sync"
	"sync/atomic"
)

var (
	counterLock sync.Mutex
	idInc       uint64
)

// EntityComponentAddedDelegate defines the delegate that gets
// called when a component gets added to the entity.
type EntityComponentAddedDelegate func(entityId uint64, componentName string)

// EntityComponentRemovedDelegate defines the delegate that gets
// called when a component gets removed from the entity.
type EntityComponentRemovedDelegate func(entityId uint64, componentName string)

// Entity defines an entity of the ecs.
type Entity struct {
	id             uint64
	components     []ComponentInterface
	addDelegate    EntityComponentAddedDelegate
	removeDelegate EntityComponentRemovedDelegate
}

// NewEntity creates a new entity with a unique identifier.
func NewEntity(componentCapacity int, add EntityComponentAddedDelegate,
	remove EntityComponentRemovedDelegate) *Entity {
	result := &Entity{
		id:             atomic.AddUint64(&idInc, 1),
		components:     make([]ComponentInterface, 0, componentCapacity),
		addDelegate:    add,
		removeDelegate: remove,
	}
	return result
}

// ID returns the id of this entity.
func (e Entity) ID() uint64 {
	return e.id
}

// Add component(s) to the entity.
func (e *Entity) Add(components ...ComponentInterface) *Entity {
	for _, component := range components {
		e.components = append(e.components, component)
		e.addDelegate(e.id, component.GetName())
	}

	return e
}

// Remove removes a component.
func (e *Entity) Remove(componentNames ...string) {
	for _, name := range componentNames {
		index := -1
		for i, c := range e.components {
			if c.GetName() == name {
				index = i
				break
			}
		}

		if index != -1 {
			copy(e.components[index:], e.components[index+1:])
			e.components[len(e.components)-1] = nil
			e.components = e.components[:len(e.components)-1]
			e.removeDelegate(e.id, name)
		}
	}
}

// Get returns a component or nil, false if not found.
func (e *Entity) Get(componentName string) (ComponentInterface, bool) {
	for _, c := range e.components {
		if c.GetName() == componentName {
			return c, true
		}
	}

	return nil, false
}

// HasComponent returns true if the entity has the component.
func (e *Entity) HasComponent(componentName string) bool {
	for _, component := range e.components {
		if component.GetName() == componentName {
			return true
		}
	}

	return false
}

// Has returns true if the entity has all the components.
func (e *Entity) Has(componentNames []string) bool {
	for _, name := range componentNames {
		if !e.HasComponent(name) {
			return false
		}
	}

	return true
}
