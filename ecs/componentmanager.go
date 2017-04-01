package ecs

import (
	"errors"
)

type componentList []ComponentInterface
type entityList []EntityID

// ComponentManager manages the components.
type ComponentManager struct {
	componentsByEntity  map[EntityID]componentList
	entitiesByComponent map[string]entityList
}

// NewComponentManager creates a new ComponentManager.
func NewComponentManager() ComponentManager {
	return ComponentManager{
		componentsByEntity:  make(map[EntityID]componentList),
		entitiesByComponent: make(map[string]entityList),
	}
}

// Add adds a component to the entity.
func (m *ComponentManager) Add(
	id EntityID, comp ComponentInterface) ComponentInterface {

	comp.SetEntityID(id)
	m.addToComponentsByEntity(id, comp)
	m.addToEntitiesByComponent(id, comp)

	return comp
}

// Get returns a ComponentInterface; and an error if it's not found.
func (m ComponentManager) Get(id EntityID, name string) (ComponentInterface, error) {
	components := m.componentsByEntity[id]
	for _, component := range components {
		if component.GetName() == name {
			return component, nil
		}
	}
	return nil, errors.New("No component by that name")
}

// GetEntities returns all the entities with the given components.
func (m ComponentManager) GetEntities(names []string) []EntityID {
	entities := make([]EntityID, 0)
	for _, name := range names {
		for _, id := range m.entitiesByComponent[name] {
			entities = append(entities, id)
		}
	}
	return entities
}

// GetEntityComponents gets all components with that name (type).
func (m ComponentManager) GetEntityComponents(
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
func (m ComponentManager) Has(id EntityID, name string) bool {
	components := m.componentsByEntity[id]
	for _, component := range components {
		if component.GetName() == name {
			return true
		}
	}
	return false
}

// HasAll checks if the Entity has all the specified Component.
func (m ComponentManager) HasAll(id EntityID, names ...string) bool {
	for _, name := range names {
		if !m.Has(id, name) {
			return false
		}
	}
	return true
}

func (m *ComponentManager) addToComponentsByEntity(
	id EntityID, comp ComponentInterface) {

	components, ok := m.componentsByEntity[id]
	if !ok {
		components = make(componentList, 0)
	}

	m.componentsByEntity[id] = append(components, comp)
}

func (m *ComponentManager) addToEntitiesByComponent(
	id EntityID, comp ComponentInterface) {

	entities, ok := m.entitiesByComponent[comp.GetName()]
	if !ok {
		entities = make(entityList, 0)
	}

	m.entitiesByComponent[comp.GetName()] = append(entities, id)
}
