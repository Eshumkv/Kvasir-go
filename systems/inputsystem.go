package systems

import "github.com/veandco/go-sdl2/sdl"

//------------------------------------------------------------------------------
// Input system

// InputSystem is the system that handles input.
type InputSystem struct {
	commands     []bool
	quitDelegate func()
}

// NewInputSystem creates a new InputSystem
func NewInputSystem(delegate func()) *InputSystem {
	return &InputSystem{
		commands:     make([]bool, 10),
		quitDelegate: delegate,
	}
}

// Update updates this system.
func (system *InputSystem) Update(dt float64) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			system.quitDelegate()
		case *sdl.KeyDownEvent:
			system.commands[toCommand(t.Keysym.Sym)] = true
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
