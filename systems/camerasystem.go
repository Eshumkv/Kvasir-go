package systems

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------------------------------------------------
// Input system

// CameraSystem is the system that handles the camera.
type CameraSystem struct {
	systemName string
	window     *sdl.Window
	halfWidth  int
	halfHeight int
	x, y       int
}

// NewCameraSystem creates a new CameraSystem
func NewCameraSystem(window *sdl.Window) *CameraSystem {
	w, h := window.GetSize()
	return &CameraSystem{
		systemName: "CameraSystem",
		window:     window,
		halfWidth:  w / 2,
		halfHeight: h / 2,
	}
}

// Update updates this system.
func (system *CameraSystem) Update(
	entities []ecs.Entity, world *ecs.World, dt float64) {

	w, h := system.window.GetSize()
	system.halfWidth = w / 2
	system.halfHeight = h / 2

	if len(entities) > 0 {
		comp, err := world.GetComponent(entities[0], "Spatial")
		if err != nil {
			return
		}
		spatial := comp.(*components.SpatialComponent)
		system.SetLocation(spatial.X, spatial.Y)
	}
}

// GetComponentNames gives a list of components that this system uses.
func (system CameraSystem) GetComponentNames() []string {
	return []string{"CameraFollow", "Spatial"}
}

// GetSystemName returns the name of this system.
func (system CameraSystem) GetSystemName() string {
	return system.systemName
}

// GetIsConcurrent checks whether this system will run in a seperate thread.
func (system CameraSystem) GetIsConcurrent() bool {
	return true
}

// GetLocation gets the current location of the camera.
func (system CameraSystem) GetLocation() (x, y int) {
	return system.x, system.y
}

// SetLocation sets the location of the camera.
func (system *CameraSystem) SetLocation(x, y int) {
	system.x, system.y = x, y
}

// GetScreenLocation calculates the correct position for the entity.
func (system *CameraSystem) GetScreenLocation(x, y int) (int, int) {
	screenX := (x - system.x) + system.halfWidth
	screenY := (y - system.y) + system.halfHeight
	return int(screenX), int(screenY)
}

// GetMousePositionInWorld gets the mouse position translated to the world
// coordinates.
func (system CameraSystem) GetMousePositionInWorld(
	mouseX, mouseY int) (int, int) {

	relativeX := system.halfWidth - mouseX
	relativeY := system.halfHeight - mouseY

	return system.x - relativeX, system.y - relativeY
}
