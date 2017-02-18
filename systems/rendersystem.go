package systems

import (
	"github.com/eshumkv/Kvasir-go/components"
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

// RenderSystem defines the system to render stuff to the screen.
type RenderSystem struct {
	renderer *sdl.Renderer
	entities []ecs.Entity
}

// NewRenderSystem returns a pointer to a new RenderSystem.
func NewRenderSystem(renderer *sdl.Renderer) *RenderSystem {
	return &RenderSystem{
		renderer: renderer,
		entities: make([]ecs.Entity, 0),
	}
}

// Init initializes the system.
func (s *RenderSystem) Init(mngr *ecs.SystemManager) {
	// Empty :(
}

// Add adds an entity to the system.
func (s *RenderSystem) Add(e *ecs.Entity) {
	s.entities = append(s.entities, *e)
}

// Update handles the update of the system.
func (s *RenderSystem) Update(dt float64) {
	for _, entity := range s.entities {
		genericComponent, ok := entity.Get("*components.ColorComponent")
		if !ok {
			continue
		}
		comp := genericComponent.(*components.ColorComponent)
		rect := entity.Rect()

		r, g, b, _, _ := s.renderer.GetDrawColor()
		s.renderer.SetDrawColor(comp.R, comp.G, comp.B, 255)
		s.renderer.FillRect(&rect)
		s.renderer.SetDrawColor(r, g, b, 255)
	}
}

// Delete deletes an entity from this system.
func (s *RenderSystem) Delete(e ecs.Entity) {
	var delete = -1
	for index, entity := range s.entities {
		if entity.ID() == e.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

// Priority defines the priority of this system.
func (s RenderSystem) Priority() uint {
	return 50
}
