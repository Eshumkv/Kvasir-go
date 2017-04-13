package components

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type PlayerComponent struct {
	name         string
	active       bool
	entity       ecs.Entity
	Speed        float64
	ShootTimeout float64
	ShootTimer   float64
	CanShoot     bool
}

func NewPlayerComponent() *PlayerComponent {
	return &PlayerComponent{
		name:         "Player",
		active:       true,
		Speed:        200,
		ShootTimeout: 0.5,
		CanShoot:     true,
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
