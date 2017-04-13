package components

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type BulletComponent struct {
	name   string
	active bool
	entity ecs.Entity
	DestX  int
	DestY  int
	Speed  float64
}

func NewBulletComponent(gotoX, gotoY int, speed float64) *BulletComponent {
	return &BulletComponent{
		name:   "Bullet",
		active: true,
		DestX:  gotoX,
		DestY:  gotoY,
		Speed:  speed,
	}
}

func (c *BulletComponent) SetActive(state bool) {
	c.active = state
}

func (c BulletComponent) GetName() string {
	return c.name
}

func (c *BulletComponent) SetEntityID(id ecs.Entity) {
	c.entity = id
}

func (c BulletComponent) GetEntityID() ecs.Entity {
	return c.entity
}
