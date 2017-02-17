package kvasir

import (
	"fmt"

	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderSystem struct {
	renderer *sdl.Renderer
	entities []ecs.Entity
}

func NewRenderSystem(renderer *sdl.Renderer) *RenderSystem {
	return &RenderSystem{
		renderer: renderer,
		entities: make([]ecs.Entity, 50),
	}
}

func (s *RenderSystem) Init(mngr *ecs.SystemManager) {
	// Empty :(
}

func (s *RenderSystem) Add(e *ecs.Entity) {
	s.entities = append(s.entities, *e)
}

func (s *RenderSystem) Update(dt float64) {
	for _, entity := range s.entities {
		comp, ok := entity.Get("ColorComponent")
		if !ok {
			continue
		}
		fmt.Println(comp)
	}
}

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

func (s RenderSystem) Priority() uint {
	return 0
}
