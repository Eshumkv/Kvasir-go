package systems

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
)

//------------------------------------------------------------------------------
// BulletSystem

// BulletSystem is the system that handles the camera.
type BulletSystem struct {
	systemName string
}

// NewBulletSystem creates a new BulletSystem
func NewBulletSystem() *BulletSystem {
	return &BulletSystem{
		systemName: "BulletSystem",
	}
}

// Update updates this system.
func (system *BulletSystem) Update(
	entities []ecs.Entity, world *ecs.World, dt float64) {

	for _, entity := range entities {
		c, err := world.GetComponent(entity, "Bullet")
		if err != nil {
			continue
		}
		bullet := c.(*components.BulletComponent)
	}
}

// GetComponentNames gives a list of components that this system uses.
func (system BulletSystem) GetComponentNames() []string {
	return []string{"Bullet"}
}

// GetSystemName returns the name of this system.
func (system BulletSystem) GetSystemName() string {
	return system.systemName
}

// GetIsConcurrent checks whether this system will run in a seperate thread.
func (system BulletSystem) GetIsConcurrent() bool {
	return true
}

func lerp(v0, v1, t float64) float64 {
	return (1-t)*v0 + t*v1
}
