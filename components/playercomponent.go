package components

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type PlayerComponent struct {
	name   string
	active bool
	entity ecs.Entity
	Speed  float64
}

func NewPlayerComponent() *PlayerComponent {
	return &PlayerComponent{
		name:   "Player",
		active: true,
		Speed:  200,
	}
}

func (c *PlayerComponent) SetActive(state bool) {
	c.active = state
}

func (c PlayerComponent) GetName() string {
	return c.name
}

func (c *PlayerComponent) SetEntityID(id ecs.Entity) {
	c.entity = id
}

func (c PlayerComponent) GetEntityID() ecs.Entity {
	return c.entity
}
