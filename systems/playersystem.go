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

type position struct{ x, y int }

// Update updates this system.
func (system *PlayerSystem) Update(
	entities []ecs.Entity, world *ecs.World, dt float64) {

	s := world.GetSystem("InputSystem")
	inputSystem := s.(*InputSystem)

	s = world.GetSystem("CameraSystem")
	camera := s.(*CameraSystem)

	for _, entity := range entities {
		t, err := world.GetComponent(entity, "Player")
		if err != nil {
			continue
		}
		playerComponent := t.(*components.PlayerComponent)

		move(*playerComponent, dt, world, inputSystem)

		if !playerComponent.CanShoot {
			playerComponent.ShootTimer += dt

			if playerComponent.ShootTimer >= playerComponent.ShootTimeout {
				playerComponent.CanShoot = true
				playerComponent.ShootTimer = 0
			}
		}

		// Shooting?
		if inputSystem.IsKeyDown(CommandShoot) &&
			playerComponent.CanShoot {

			screenMouseX, screenMouseY := inputSystem.GetMousePosition()
			mx, my := camera.GetMousePositionInWorld(screenMouseX, screenMouseY)

			// Shoot the bullet
			entity := world.Create()
			world.AddComponents(entity,
				components.NewRenderComponent(255, 255, 255),
				components.NewSpatialComponent(mx, my, 1, 10, 10),
				components.NewBulletComponent(mx, my))
			playerComponent.CanShoot = false
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

// GetIsConcurrent checks whether this system will run in a seperate thread.
func (system PlayerSystem) GetIsConcurrent() bool {
	return true
}

func move(playerComponent components.PlayerComponent,
	dt float64, world *ecs.World, inputSystem *InputSystem) {

	x := 0.0
	y := 0.0

	if inputSystem.IsKeyDown(CommandLeft) {
		x = -playerComponent.Speed * dt
	}
	if inputSystem.IsKeyDown(CommandRight) {
		x = playerComponent.Speed * dt
	}

	if inputSystem.IsKeyDown(CommandUp) {
		y = -playerComponent.Speed * dt
	}
	if inputSystem.IsKeyDown(CommandDown) {
		y = playerComponent.Speed * dt
	}

	if x != 0 || y != 0 {
		c, err := world.GetComponent(playerComponent.GetEntityID(), "Spatial")
		if err != nil {
			return
		}
		spatial := c.(*components.SpatialComponent)

		spatial.X = int(float64(spatial.X) + x)
		spatial.Y = int(float64(spatial.Y) + y)
	}
}
