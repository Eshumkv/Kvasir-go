package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
)

type EntityManagerItems struct {
	systems map[string]bool
	entity  *ecs.Entity
}

// MyEntityManager implements the ecs/EntityManager interface.
type MyEntityManager struct {
	items map[uint64]EntityManagerItems
	mngr  *ecs.SystemManager
}

// NewMyEntityManager returns a new MyEntityManager.
func NewMyEntityManager() *MyEntityManager {
	return &MyEntityManager{
		items: make(map[uint64]EntityManagerItems, 0),
	}
}

// Add adds an entity to this manager.
func (s *MyEntityManager) Add(entity ecs.Entity, systemIDs ...string) {
	systemMap := make(map[string]bool)
	for _, system := range systemIDs {
		systemMap[system] = true
	}
	s.items[entity.ID()] = EntityManagerItems{
		systems: systemMap,
		entity:  &entity,
	}
}

// Delete removes an entity from this manager. if systemIDs is nil, delete entire entity, else only the systems defined.
func (s *MyEntityManager) Delete(id uint64, systemIDs ...string) {
	if systemIDs == nil {
		delete(s.items, id)
	} else {
		item := s.items[id]
		newMap := make(map[string]bool)

		for system := range item.systems {
			if !contains(systemIDs, system) {
				newMap[system] = true
			}
		}

		item.systems = newMap
	}
}

func contains(list []string, item string) bool {
	return true
}

// Get returns an entity managed by this manager.
func (s MyEntityManager) Get(id uint64) ecs.Entity {
	return *s.items[id].entity
}

// Init initializes the system.
func (s *MyEntityManager) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Update handles the update of the system.
func (s *MyEntityManager) Update(dt float64) {
}

// Priority defines the priority of this system.
func (s MyEntityManager) Priority() uint {
	return 50
}

// HandleMessage handles any messages that need to be dealt with.
func (s *MyEntityManager) HandleMessage(
	msg ecs.Message, data interface{}) interface{} {
	switch msg {
	case MessageGetEntitiesOfSystem:
		systemID := data.(string)
		entities := make([]ecs.Entity, 0)
		for _, item := range s.items {
			if item.systems[systemID] {
				entities = append(entities, *item.entity)
			}
		}

		if len(entities) == 0 {
			return nil
		}

		return entities
	case MessageSetEntityLocation:
		locationData := data.(struct {
			id     uint64
			system string
			x      float64
			y      float64
		})

		// This system is not allowed to edit this entity
		if !s.items[locationData.id].systems[locationData.system] {
			return nil
		}

		entity := s.items[locationData.id].entity

		entity.SetX(locationData.x)
		entity.SetY(locationData.y)
	}

	return nil
}
