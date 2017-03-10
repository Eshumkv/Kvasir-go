package ecs

// EntityManager defines the interface for an entitymanager.
type EntityManager interface {
	Add(entity Entity, systemIDs ...string)
	Delete(id uint64, systemIDs ...string)
	Get(id uint64) Entity
	HandleMessage(msg Message, data interface{}) interface{}
}
