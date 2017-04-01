package systems

import "github.com/veandco/go-sdl2/sdl"
import "github.com/Eshumkv/kvasir-go/ecs"

//------------------------------------------------------------------------------
// Input system

// InputSystem is the system that handles input.
type InputSystem struct {
	commands           []bool
	quitDelegate       func()
	fullscreenDelegate func()
	systemName         string
}

// NewInputSystem creates a new InputSystem
func NewInputSystem(delegate func(), fullscreen func()) *InputSystem {
	return &InputSystem{
		commands:           make([]bool, 10),
		quitDelegate:       delegate,
		fullscreenDelegate: fullscreen,
		systemName:         "InputSystem",
	}
}

// Update updates this system.
func (system *InputSystem) Update(
	entities []ecs.EntityID, world *ecs.World, dt float64) {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			system.quitDelegate()
		case *sdl.KeyDownEvent:
			system.commands[toCommand(t.Keysym.Sym)] = true

			if system.commands[CommandFullscreen] {
				system.fullscreenDelegate()
				system.commands[CommandFullscreen] = false
			}
		case *sdl.KeyUpEvent:
			system.commands[toCommand(t.Keysym.Sym)] = false
		}
	}
}

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

// GetComponentNames gives a list of components that this system uses.
func (system InputSystem) GetComponentNames() []string {
	var myComponents []string
	return myComponents
}

// GetSystemName returns the name of this system.
func (system InputSystem) GetSystemName() string {
	return system.systemName
}

// IsKeyDown checks if the specified key is down.
func (system InputSystem) IsKeyDown(command Command) bool {
	return system.commands[command]
}

//------------------------------------------------------------------------------
// Commands

// Command defines the type used for the Command enum
type Command int

// The Command enum
const (
	CommandNone Command = iota
	CommandFullscreen
	CommandShoot
	CommandLeft
	CommandUp
	CommandRight
	CommandDown
	CommandSpeedup
	CommandCount
)
