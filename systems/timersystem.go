package systems

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
)

//------------------------------------------------------------------------------
// TimerSystem

// TimerSystem is the system that handles the camera.
type TimerSystem struct {
	systemName string
}

// NewTimerSystem creates a new TimerSystem
func NewTimerSystem() *TimerSystem {
	return &TimerSystem{
		systemName: "TimerSystem",
	}
}

// Update updates this system.
func (system *TimerSystem) Update(
	entities []ecs.Entity, world *ecs.World, dt float64) {

	for _, entity := range entities {
		c, err := world.GetComponent(entity, "Timer")
		if err != nil {
			continue
		}
		timer := c.(*components.TimerComponent)

		if timer.JustStarted {
			timer.Timer = 0
			timer.IsStarted = true
			timer.JustStarted = false
		}

		if timer.IsStarted {
			timer.Timer += dt
			if timer.Timer >= timer.Time {
				timer.IsStarted = false
				timer.JustStarted = false
				timer.Callback(timer, world)
			}
		}
	}
}

// GetComponentNames gives a list of components that this system uses.
func (system TimerSystem) GetComponentNames() []string {
	return []string{"Timer"}
}

// GetSystemName returns the name of this system.
func (system TimerSystem) GetSystemName() string {
	return system.systemName
}

// GetIsConcurrent checks whether this system will run in a seperate thread.
func (system TimerSystem) GetIsConcurrent() bool {
	return true
}
