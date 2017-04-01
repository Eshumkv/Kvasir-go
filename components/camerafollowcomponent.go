package components

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type CameraFollowComponent struct {
	name   string
	active bool
	entity ecs.EntityID
}

func NewCameraFollowComponent() *CameraFollowComponent {
	return &CameraFollowComponent{
		name:   "CameraFollow",
		active: true,
	}
}

func (c *CameraFollowComponent) SetActive(state bool) {
	c.active = state
}

func (c CameraFollowComponent) GetName() string {
	return c.name
}

func (c *CameraFollowComponent) SetEntityID(id ecs.EntityID) {
	c.entity = id
}

func (c CameraFollowComponent) GetEntityID() ecs.EntityID {
	return c.entity
}
