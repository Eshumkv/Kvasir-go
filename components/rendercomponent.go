package components

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type RenderComponent struct {
	name    string
	active  bool
	entity  ecs.Entity
	R, G, B uint8
}

func NewRenderComponent(r, g, b uint8) *RenderComponent {
	return &RenderComponent{
		name:   "Render",
		active: true,
		R:      r,
		G:      g,
		B:      b,
	}
}

func (c *RenderComponent) SetActive(state bool) {
	c.active = state
}

func (c RenderComponent) GetName() string {
	return c.name
}

func (c *RenderComponent) SetEntityID(id ecs.Entity) {
	c.entity = id
}

func (c RenderComponent) GetEntityID() ecs.Entity {
	return c.entity
}
