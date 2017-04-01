package ecs

// EntityManager manages the instances of the entities.
type EntityManager struct {
	entities []Entity
}

// NewEntityManager returns a new instance of the EntityManager.
func NewEntityManager() EntityManager {
	return EntityManager{
		entities: make([]Entity, 0),
	}
}
