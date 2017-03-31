package systems

import (
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
	return &RenderSystem{
		renderer:   renderer,
		systemName: "RenderSystem",
	}
}

// Update updates this system this frame.
func (system *RenderSystem) Update(entities []*ecs.Entity, dt float64) {
	system.renderer.Clear()

	// Draw here

	system.renderer.Present()
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
