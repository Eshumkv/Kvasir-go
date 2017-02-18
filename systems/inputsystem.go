package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

// InputSystem defines the system to process input.
type InputSystem struct {
	game     GameInterface
	entities []ecs.Entity
}

// NewInputSystem returns a pointer to a new InputSystem.
func NewInputSystem(game GameInterface) *InputSystem {
	return &InputSystem{
		game:     game,
		entities: make([]ecs.Entity, 0),
	}
}

// Init initializes the system.
func (s *InputSystem) Init(mngr *ecs.SystemManager) {
	// Empty :(
}

// Add adds an entity to the system.
func (s *InputSystem) Add(e *ecs.Entity) {
	s.entities = append(s.entities, *e)
}

// Update handles the update of the system.
func (s *InputSystem) Update(dt float64) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			s.game.SetRunning(false)
		case *sdl.KeyDownEvent:
			s.game.SetCommand(toCommand(t.Keysym.Sym), true)

			if s.game.GetCommand(CommandFullscreen) {
				s.game.ToggleFullscreen()
				s.game.SetCommand(CommandFullscreen, false)
			}

		case *sdl.KeyUpEvent:
			s.game.SetCommand(toCommand(t.Keysym.Sym), false)
		}
	}
}

// toCommand turns an SDL2 keycode into a game command.
func toCommand(keycode sdl.Keycode) Command {
	switch keycode {
	case sdl.K_F11:
		return CommandFullscreen
	case sdl.K_w:
		return CommandUp
	case sdl.K_s:
		return CommandDown
	case sdl.K_a:
		return CommandLeft
	case sdl.K_d:
		return CommandRight
	case sdl.K_SPACE:
		return CommandShoot
	case sdl.K_LSHIFT:
		return CommandSpeedup
	default:
		return CommandNone
	}
}

// Delete deletes an entity from this system.
func (s *InputSystem) Delete(e ecs.Entity) {
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
func (s InputSystem) Priority() uint {
	return 0
}
