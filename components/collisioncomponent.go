package components

// CollisionComponent defines a component that can be used to handle collisions.
type CollisionComponent struct {
	location uint64
}

// NewCollisionComponent returns a pointer to a new CollisionComponent.
func NewCollisionComponent() *CollisionComponent {
	return &CollisionComponent{
		location: 0,
	}
}
