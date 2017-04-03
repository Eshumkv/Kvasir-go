package ecs

import (
	"sync"
	"sync/atomic"
)

var (
	counterLock sync.Mutex
	idInc       uint32
)

// World defines the interface for the Entity Component System.
type World struct {
	em           EntityManager
	dt           float64
	systems      []SystemInterface
	systemsCache map[string]int
}

// NewWorld returns a new World object.
func NewWorld(allSystems []SystemInterface) World {
	return World{
		em:           NewEntityManager(),
		systems:      allSystems,
		systemsCache: make(map[string]int),
	}
}

// Create a new Entity.
func (world *World) Create() Entity {
	id := atomic.AddUint32(&idInc, 1)
	entity := Entity(id)
	world.em.Add(entity)

	return entity
}

// Remove an entity from the world.
func (world *World) Remove(id Entity) {
	world.em.Remove(id)
}

// ClearEntities removes all entities from the world.
func (world *World) ClearEntities() {
	world.em.RemoveAllEntities()
}

// SetDeltaTime sets the delta time.
func (world *World) SetDeltaTime(dt float64) {
	world.dt = dt
}

// Update the world state.
func (world *World) Update() {
	for _, system := range world.systems {
		system.Update(
			world.em.GetEntities(system.GetComponentNames()),
			world,
			world.dt)
	}
	world.em.Process()
}

// GetSystem gets a system with a specific system.
func (world *World) GetSystem(name string) SystemInterface {
	// Test the cache
	if index, ok := world.systemsCache[name]; ok {
		return world.systems[index]
	}

	for index, system := range world.systems {
		if system.GetSystemName() == name {
			world.systemsCache[name] = index
			return system
		}
	}
	panic("No such system!")
}

// AddComponents adds components to an entity.
func (world *World) AddComponents(
	id Entity, components ...ComponentInterface) {

	world.em.AddComponents(id, components...)
}

// GetComponent gets a component from an entity.
func (world World) GetComponent(
	id Entity, name string) (ComponentInterface, error) {
	return world.em.GetComponent(id, name)
}

// GetEntitiesByComponent gets all entities with that componentname (type).
func (world World) GetEntitiesByComponent(
	name string) []ComponentInterface {

	return world.em.GetEntitiesByComponent(name)
}
