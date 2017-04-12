package ecs

import (
	"errors"
)

type componentList []ComponentInterface
type entityList []Entity

// ComponentManager manages the components.
type ComponentManager struct {
	componentsByEntity  map[Entity]componentList
	entitiesByComponent map[string]entityList
}

// NewComponentManager creates a new ComponentManager.
func NewComponentManager() ComponentManager {
	return ComponentManager{
		componentsByEntity:  make(map[Entity]componentList),
		entitiesByComponent: make(map[string]entityList),
	}
}

// Add adds a component to the entity.
func (m *ComponentManager) Add(
	id Entity, comp ComponentInterface) ComponentInterface {

	comp.SetEntityID(id)
	m.addToComponentsByEntity(id, comp)
	m.addToEntitiesByComponent(id, comp)

	return comp
}

// Get returns a ComponentInterface; and an error if it's not found.
func (m ComponentManager) Get(id Entity, name string) (ComponentInterface, error) {
	components := m.componentsByEntity[id]
	for _, component := range components {
		if component.GetName() == name {
			return component, nil
		}
	}
	return nil, errors.New("No component by that name")
}

// GetEntities returns all the entities with the given components.
func (m ComponentManager) GetEntities(names []string) []Entity {
	entities := make([]Entity, 0)
	for _, name := range names {
		for _, id := range m.entitiesByComponent[name] {
			entities = append(entities, id)
		}
	}
	return entities
}

// GetEntitiesByComponent gets all entities with that componentname (type).
func (m ComponentManager) GetEntitiesByComponent(
	name string) []ComponentInterface {

	compEntities := m.GetEntities([]string{name})
	entities := make(componentList, 0)
	for _, entity := range compEntities {
		if value, err := m.Get(entity, name); err == nil {
			entities = append(entities, value)
		}
	}
	return entities
}

// Has checks whether the entity has the given components.
func (m ComponentManager) Has(id Entity, name string) bool {
	components := m.componentsByEntity[id]
	for _, component := range components {
		if component.GetName() == name {
			return true
		}
	}
	return false
}

// HasAll checks if the Entity has all the specified Component.
func (m ComponentManager) HasAll(id Entity, names ...string) bool {
	for _, name := range names {
		if !m.Has(id, name) {
			return false
		}
	}
	return true
}

// DeleteEntity removes an entity and all it's components
func (m *ComponentManager) DeleteEntity(id Entity) {
	components := m.componentsByEntity[id]

	delete(m.componentsByEntity, id)
	for _, component := range components {
		entities := m.entitiesByComponent[component.GetName()]
		index := -1
		for i, entity := range entities {
			if entity == id {
				index = i
				break
			}
		}
		if index != -1 {
			m.entitiesByComponent[component.GetName()] =
				append(entities[:index], entities[index+1:]...)
		}
	}
}

func (m *ComponentManager) addToComponentsByEntity(
	id Entity, comp ComponentInterface) {

	components, ok := m.componentsByEntity[id]
	if !ok {
		components = make(componentList, 0)
	}

	m.componentsByEntity[id] = append(components, comp)
}

func (m *ComponentManager) addToEntitiesByComponent(
	id Entity, comp ComponentInterface) {

	entities, ok := m.entitiesByComponent[comp.GetName()]
	if !ok {
		entities = make(entityList, 0)
	}

	m.entitiesByComponent[comp.GetName()] = append(entities, id)
}
