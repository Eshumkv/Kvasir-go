package systems

import "github.com/veandco/go-sdl2/sdl"

//------------------------------------------------------------------------------
// Render system

type RenderSystem struct {
	renderer *sdl.Renderer
}

func NewRenderSystem(renderer *sdl.Renderer) *RenderSystem {
	return &RenderSystem{
		renderer: renderer,
	}
}

func (system *RenderSystem) Update(dt float64) {
	system.renderer.Clear()

	// Draw here

	system.renderer.Present()
}
