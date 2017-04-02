package ecs

// ComponentInterface defines the interface for a component.
type ComponentInterface interface {
	SetActive(state bool)
	GetName() string
	GetEntityID() Entity
	SetEntityID(id Entity)
}
