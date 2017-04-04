package systems

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
)

//------------------------------------------------------------------------------
// PlayerSystem

// PlayerSystem is the system that handles the camera.
type PlayerSystem struct {
	systemName string
}

// NewPlayerSystem creates a new PlayerSystem
func NewPlayerSystem() *PlayerSystem {
	return &PlayerSystem{
		systemName: "PlayerSystem",
	}
}

// Update updates this system.
func (system *PlayerSystem) Update(
	entities []ecs.Entity, world *ecs.World, dt float64) {

	players := world.GetEntitiesByComponent("Player")
	s := world.GetSystem("InputSystem")
	inputSystem := s.(*InputSystem)

	for _, comp := range players {
		player := comp.(*components.PlayerComponent)
		x := 0.0
		y := 0.0

		if inputSystem.IsKeyDown(CommandLeft) {
			x = -player.Speed * dt
		}
		if inputSystem.IsKeyDown(CommandRight) {
			x = player.Speed * dt
		}

		if inputSystem.IsKeyDown(CommandUp) {
			y = -player.Speed * dt
		}
		if inputSystem.IsKeyDown(CommandDown) {
			y = player.Speed * dt
		}

		if x != 0 || y != 0 {
			c, err := world.GetComponent(player.GetEntityID(), "Spatial")
			if err != nil {
				continue
			}
			spatial := c.(*components.SpatialComponent)

			spatial.X = int(float64(spatial.X) + x)
			spatial.Y = int(float64(spatial.Y) + y)
		}
	}
}

// GetComponentNames gives a list of components that this system uses.
func (system PlayerSystem) GetComponentNames() []string {
	return []string{"Player", "Spatial"}
}

// GetSystemName returns the name of this system.
func (system PlayerSystem) GetSystemName() string {
	return system.systemName
}
