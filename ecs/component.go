package ecs

// Component specifies the interface that every component should follow.
type Component interface {
}

// ComponentHandler specifies what objects can handle a list of components.
type ComponentHandler interface {
	Add(comp Component)
	Remove(comp Component)
	Get(comp Component) (Component, bool)
	Has(comp Component) bool
}
