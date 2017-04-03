package ecs

// EntityManager manages the instances of the entities.
type EntityManager struct {
	entities map[Entity]bool
	added    map[Entity]bool
	removed  map[Entity]bool
	cm       ComponentManager
}

// NewEntityManager returns a new instance of the EntityManager.
func NewEntityManager() EntityManager {
	return EntityManager{
		entities: make(map[Entity]bool),
		added:    make(map[Entity]bool),
		removed:  make(map[Entity]bool),
		cm:       NewComponentManager(),
	}
}

// Add a new Entity to the manager.
func (em *EntityManager) Add(id Entity) {
	em.added[id] = true
}

// Remove an entity from the manager.
func (em *EntityManager) Remove(id Entity) {
	em.removed[id] = true
}

// RemoveAllEntities removes all entities.
func (em *EntityManager) RemoveAllEntities() {
	for entity := range em.entities {
		em.Remove(entity)
	}
}

// Process handles all entity updates.
func (em *EntityManager) Process() {
	for entity := range em.removed {
		delete(em.entities, entity)
		em.cm.DeleteEntity(entity)
	}
	for entity := range em.added {
		em.entities[entity] = true
	}
	em.removed = make(map[Entity]bool)
	em.added = make(map[Entity]bool)
}

// AddComponents adds components to an entity.
func (em *EntityManager) AddComponents(
	id Entity, components ...ComponentInterface) {

	for _, component := range components {
		em.cm.Add(id, component)
	}
}

// GetEntities returns all the entities with the given components.
func (em EntityManager) GetEntities(names []string) []Entity {
	return em.cm.GetEntities(names)
}

// GetComponent gets a component from an entity.
func (em EntityManager) GetComponent(
	id Entity, name string) (ComponentInterface, error) {

	return em.cm.Get(id, name)
}

// GetEntitiesByComponent gets all entities with that componentname (type).
func (em EntityManager) GetEntitiesByComponent(
	name string) []ComponentInterface {

	return em.cm.GetEntitiesByComponent(name)
}
