package components

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type TimerComponentCallback func(timer *TimerComponent, world *ecs.World)

type TimerComponent struct {
	name   string
	active bool
	entity ecs.Entity

	// Data
	Timer       float64
	Time        float64
	IsStarted   bool
	JustStarted bool
	Callback    TimerComponentCallback
}

func NewTimerComponent(
	time float64, cb TimerComponentCallback, start bool) *TimerComponent {

	return &TimerComponent{
		name:        "Timer",
		active:      true,
		Time:        time,
		JustStarted: start,
		IsStarted:   start,
		Callback:    cb,
	}
}

func (c *TimerComponent) SetActive(state bool) {
	c.active = state
}

func (c TimerComponent) GetName() string {
	return c.name
}

func (c *TimerComponent) SetEntityID(id ecs.Entity) {
	c.entity = id
}

func (c TimerComponent) GetEntityID() ecs.Entity {
	return c.entity
}

func (c *TimerComponent) Start() {
	c.JustStarted = true
	c.IsStarted = true
}
