package ecs

// EntityManager manages the instances of the entities.
type EntityManager struct {
	entities []Entity
	added    []Entity
	changed  []Entity
	removed  []Entity
}

// NewEntityManager returns a new instance of the EntityManager.
func NewEntityManager() EntityManager {
	return EntityManager{
		entities: make([]Entity, 0),
		added:    make([]Entity, 0, 10),
		changed:  make([]Entity, 0, 10),
		removed:  make([]Entity, 0, 10),
	}
}

func (em *EntityManager) Add(id Entity) {
	em.added = append(em.added, id)
}

func (em *EntityManager) Remove(id Entity) {
	em.removed = append(em.removed, id)
}
