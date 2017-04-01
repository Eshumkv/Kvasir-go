package ecs

// EntityID defines an ID of an entity.
type EntityID uint32

// Entity defines an entity of the ecs. Should only be gotten from
// World.Create()
type Entity struct {
	id         EntityID
	components []ComponentInterface
}

// ID returns the id of this entity.
func (e Entity) ID() EntityID {
	return e.id
}

// Add component(s) to the entity.
func (e *Entity) Add(components ...ComponentInterface) {
	for _, component := range components {
		e.components = append(e.components, component)
	}

	//return e
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
