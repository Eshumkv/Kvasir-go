package ecs

// SystemInterface defines the contract for all systems to follow.
type SystemInterface interface {
	Update(entities []Entity, world *World, dt float64)
	GetComponentNames() []string
	GetSystemName() string
}
