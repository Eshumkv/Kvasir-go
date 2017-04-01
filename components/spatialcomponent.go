package components

import "github.com/Eshumkv/kvasir-go/ecs"
import "github.com/veandco/go-sdl2/sdl"

// SpatialComponent defines the location of an entity.
type SpatialComponent struct {
	name    string
	active  bool
	entity  ecs.EntityID
	X, Y, Z int
	W, H    int
}

// NewSpatialComponent creates a new SpatialComponent.
func NewSpatialComponent(x, y, z, w, h int) *SpatialComponent {
	return &SpatialComponent{
		name:   "Spatial",
		active: true,
		X:      x,
		Y:      y,
		Z:      z,
		W:      w,
		H:      h,
	}
}

// SetActive sets the state of the component.
func (c *SpatialComponent) SetActive(state bool) {
	c.active = state
}

// GetName gets the name of this component.
func (c SpatialComponent) GetName() string {
	return c.name
}

// SetEntityID sets the EntityID.
func (c *SpatialComponent) SetEntityID(id ecs.EntityID) {
	c.entity = id
}

// GetEntityID gets the EntityID.
func (c SpatialComponent) GetEntityID() ecs.EntityID {
	return c.entity
}

// GetRect returns an sdl.Rect.
func (c SpatialComponent) GetRect() sdl.Rect {
	return sdl.Rect{
		X: int32(c.X),
		Y: int32(c.Y),
		W: int32(c.W),
		H: int32(c.H),
	}
}
