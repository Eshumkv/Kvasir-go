package kvasir

import (
	"fmt"

	"github.com/eshumkv/Kvasir-go/components"
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/eshumkv/Kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

// Game struct that holds all relevant data
type Game struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	IsRunning    bool
	commands     []bool
	isFullscreen bool
	sManager     *ecs.SystemManager
	rManager     *ecs.SystemManager
}

// NewGame creates a new game.
func NewGame(window *sdl.Window, renderer *sdl.Renderer) *Game {
	result := Game{
		window:    window,
		renderer:  renderer,
		IsRunning: true,
		commands:  make([]bool, CommandCount),
		sManager:  ecs.NewSystemManager(),
		rManager:  ecs.NewSystemManager(),
	}

	// Init
	renderer.SetDrawColor(110, 132, 174, 255)

	result.setupSystems()

	return &result
}

// ProcessInput processes the input for the game.
func (game *Game) ProcessInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			game.IsRunning = false
		case *sdl.KeyDownEvent:
			game.commands[toCommand(t.Keysym.Sym)] = true

			if game.commands[CommandFullscreen] {
				var flag uint32 = sdl.WINDOW_FULLSCREEN_DESKTOP

				if game.isFullscreen {
					flag = 0
				}

				game.isFullscreen = !game.isFullscreen
				game.window.SetFullscreen(flag)
				game.commands[CommandFullscreen] = false
			}

		case *sdl.KeyUpEvent:
			game.commands[toCommand(t.Keysym.Sym)] = false
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

// Update updates the gamestate.
func (game *Game) Update(dt float64) {
	if game.commands[CommandLeft] {
		fmt.Println("Yay")
	}
	game.sManager.Update(dt)
}

// Render shows the gamestate on screen.
func (game *Game) Render(lag float64) {
	game.renderer.Clear()
	game.rManager.Update(lag)
	game.renderer.Present()
}

func (game *Game) setupSystems() {
	// Setup the render systems
	game.rManager.AddSystem(systems.NewRenderSystem(game.renderer))

	for _, system := range game.rManager.Systems() {
		switch system.(type) {
		case *systems.RenderSystem:
			entity := ecs.NewEntity(0, 0, 50, 50).Add(components.NewColorComponent(50, 60, 200))
			system.Add(entity)
		}
	}
}
