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
	cm           ComponentManager
	dt           float64
	systems      []SystemInterface
	systemsCache map[string]int
}

// NewWorld returns a new World object.
func NewWorld(allSystems []SystemInterface) World {
	return World{
		em:           NewEntityManager(),
		cm:           NewComponentManager(),
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

func (world *World) ClearEntities() {
}

// SetDeltaTime sets the delta time.
func (world *World) SetDeltaTime(dt float64) {
	world.dt = dt
}

// Update the world state.
func (world *World) Update() {
	for _, system := range world.systems {
		system.Update(
			world.cm.GetEntities(system.GetComponentNames()),
			world,
			world.dt)
	}
}

// AddComponents adds components to an entity.
func (world *World) AddComponents(
	id Entity, components ...ComponentInterface) {

	for _, component := range components {
		world.cm.Add(id, component)
	}
}

// GetComponentManager returns the componentmanager.
func (world *World) GetComponentManager() *ComponentManager {
	return &world.cm
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
