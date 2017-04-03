package systems

import (
	"sort"

	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------------------------------------------------
// Render system

// RenderSystem defines the system used for rendering the entities.
type RenderSystem struct {
	renderer   *sdl.Renderer
	systemName string
}

// NewRenderSystem initializes a new RenderSystem.
func NewRenderSystem(renderer *sdl.Renderer) *RenderSystem {
	renderer.SetDrawColor(110, 132, 174, 255)
	return &RenderSystem{
		renderer:   renderer,
		systemName: "RenderSystem",
	}
}

// Update updates this system this frame.
func (system *RenderSystem) Update(
	entities []ecs.Entity, world *ecs.World, dt float64) {

	s := world.GetSystem("CameraSystem")
	cameraSystem := s.(*CameraSystem)

	spatialEntities := system.getAndSortEntities(world)

	system.renderer.Clear()
	// Draw here
	for _, spatial := range spatialEntities {
		entityID := spatial.GetEntityID()

		comp, err := world.GetComponent(entityID, "Render")
		if err != nil {
			continue
		}
		render := comp.(*components.RenderComponent)

		sx, sy := cameraSystem.GetScreenLocation(spatial.X, spatial.Y)
		toDraw := sdl.Rect{
			X: int32(sx),
			Y: int32(sy),
			W: int32(spatial.W),
			H: int32(spatial.H)}
		r, g, b, _, _ := system.renderer.GetDrawColor()

		system.renderer.SetDrawColor(render.R, render.G, render.B, 255)
		system.renderer.FillRect(&toDraw)
		system.renderer.SetDrawColor(r, g, b, 255)
	}

	system.renderer.Present()
}

func (system RenderSystem) getAndSortEntities(
	world *ecs.World) []*components.SpatialComponent {

	list := world.GetEntitiesByComponent("Spatial")
	zEntities := make([]*components.SpatialComponent, 0, len(list))
	for _, item := range list {
		c := item.(*components.SpatialComponent)
		zEntities = append(zEntities, c)
	}
	sort.Sort(ByZ(zEntities))

	return zEntities
}

// GetComponentNames returns a list of components this system needs.
func (system RenderSystem) GetComponentNames() []string {
	return []string{
		"Spatial",
		"Render"}
}

// GetSystemName returns the name of this system.
func (system RenderSystem) GetSystemName() string {
	return system.systemName
}

// ByZ implements the sort interface for []ecs.Entity on Z value.
type ByZ []*components.SpatialComponent

func (z ByZ) Len() int           { return len(z) }
func (z ByZ) Swap(i, j int)      { z[i], z[j] = z[j], z[i] }
func (z ByZ) Less(i, j int) bool { return z[i].Z < z[j].Z }
