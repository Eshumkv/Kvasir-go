package systems

import (
	"errors"

	"sort"

	"github.com/eshumkv/Kvasir-go/components"
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/eshumkv/Kvasir-go/parts"
	"github.com/veandco/go-sdl2/sdl"
)

// RenderSystem defines the system to render stuff to the screen.
type RenderSystem struct {
	id       string
	mngr     *ecs.SystemManager
	renderer *sdl.Renderer
	camera   parts.CameraInterface
}

// NewRenderSystem returns a pointer to a new RenderSystem.
func NewRenderSystem(
	renderer *sdl.Renderer, camera parts.CameraInterface) *RenderSystem {
	return &RenderSystem{
		id:       "RenderSystem",
		renderer: renderer,
		camera:   camera,
	}
}

// Init initializes the system.
func (s *RenderSystem) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Update handles the update of the system.
func (s *RenderSystem) Update(dt float64) {
	ret, err := s.mngr.SendMessage(MessageGetEntitiesOfSystem, s.id)
	if err != nil {
		return
	}
	entities := ret.([]ecs.Entity)
	sort.Sort(ByZ(entities))
	for _, entity := range entities {
		comp, err := getColorComponent(&entity)
		if err != nil {
			continue
		}
		rect := entity.Rect()
		sx, sy := s.camera.GetScreenLocation(float64(rect.X), float64(rect.Y))
		toDraw := sdl.Rect{
			X: int32(sx),
			Y: int32(sy),
			W: entity.W(),
			H: entity.H()}

		r, g, b, _, _ := s.renderer.GetDrawColor()
		s.renderer.SetDrawColor(comp.R, comp.G, comp.B, 255)
		s.renderer.FillRect(&toDraw)
		s.renderer.SetDrawColor(r, g, b, 255)
	}
}

// Priority defines the priority of this system.
func (s RenderSystem) Priority() uint {
	return 50
}

// HandleMessage handles any messages that need to be dealt with.
func (s *RenderSystem) HandleMessage(
	msg ecs.Message, data interface{}) interface{} {
	switch msg {
	case MessageCameraUpdate:
		s.camera = data.(parts.CameraInterface)
	}
	return nil
}

func getColorComponent(e *ecs.Entity) (*components.ColorComponent, error) {
	genericComponent, ok := e.Get("*components.ColorComponent")
	if !ok {
		return nil, errors.New("No ColorComponent")
	}
	return genericComponent.(*components.ColorComponent), nil
}

// ByZ implements the sort interface for []ecs.Entity on Z value.
type ByZ []ecs.Entity

func (z ByZ) Len() int           { return len(z) }
func (z ByZ) Swap(i, j int)      { z[i], z[j] = z[j], z[i] }
func (z ByZ) Less(i, j int) bool { return z[i].Z() < z[j].Z() }
