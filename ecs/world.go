package ecs

import (
	"sync"
	"sync/atomic"

	"github.com/eshumkv/Kvasir-go/utils"
)

var (
	counterLock sync.Mutex
	idInc       uint32
)

// World defines the interface for the Entity Component System.
type World struct {
	em              EntityManager
	dt              float64
	systems         []SystemInterface
	systemsToUpdate []SystemInterface
	firsRun         bool
	systemsCache    map[string]int

	// Fps stuff
	lastFps     float64
	fps         float64
	fpsSmooting float64
}

// NewWorld returns a new World object.
func NewWorld(allSystems []SystemInterface) World {
	world := World{
		em:              NewEntityManager(),
		systems:         allSystems,
		systemsToUpdate: make([]SystemInterface, 0),
		systemsCache:    make(map[string]int),
		firsRun:         false,
		fpsSmooting:     0.9,
	}

	for _, system := range allSystems {
		if !system.GetIsConcurrent() {
			world.systemsToUpdate = append(world.systemsToUpdate, system)
		}
	}

	return world
}

// Create a new Entity.
func (world *World) Create() Entity {
	entity := Entity(atomic.AddUint32(&idInc, 1))
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
	// if world.firsRun {
	// 	world.firsRun = false

	// 	for _, system := range world.systems {
	// 		if system.GetIsConcurrent() {
	// 			go system.Update(
	// 				world.em.GetEntities(system.GetComponentNames()),
	// 				world,
	// 				world.dt)
	// 		}
	// 	}
	// }

	for _, system := range world.systems {
		system.Update(
			world.em.GetEntities(system.GetComponentNames()),
			world,
			world.dt)
	}
	world.em.Process()

	world.lastFps = world.fps
	world.fps = (world.lastFps * world.fpsSmooting) + ((1.0 / world.dt) * (1.0 - world.fpsSmooting))
	utils.DEBUG(world.fps)
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
	panic("No such system with name " + name + "!")
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

// FPS returns the current frames per second
func (world World) FPS() float64 {
	return world.fps
}
