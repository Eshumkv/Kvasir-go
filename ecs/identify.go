package ecs

// Identifier specifies something that can be uniquely identified.
type Identifier interface {
	ID() uint64
}
