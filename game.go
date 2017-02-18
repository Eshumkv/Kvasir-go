package kvasir

import (
	"github.com/eshumkv/Kvasir-go/components"
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/eshumkv/Kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

// Game struct that holds all relevant data
type Game struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	isRunning    bool
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
		isRunning: true,
		commands:  make([]bool, systems.CommandCount),
		sManager:  ecs.NewSystemManager(),
		rManager:  ecs.NewSystemManager(),
	}

	// Init
	renderer.SetDrawColor(110, 132, 174, 255)

	result.setupSystems()

	return &result
}

// Update updates the gamestate.
func (game *Game) Update(dt float64) {
	game.sManager.Update(dt)
}

// Render shows the gamestate on screen.
func (game *Game) Render(lag float64) {
	game.renderer.Clear()
	game.rManager.Update(lag)
	game.renderer.Present()
}

func (game *Game) IsRunning() bool {
	return game.isRunning
}

func (game *Game) SetRunning(state bool) {
	game.isRunning = state
}

func (game *Game) SetCommand(c systems.Command, state bool) {
	game.commands[c] = state
}

func (game *Game) GetCommand(c systems.Command) bool {
	return game.commands[c]
}

func (game *Game) ToggleFullscreen() {
	var flag uint32 = sdl.WINDOW_FULLSCREEN_DESKTOP

	if game.isFullscreen {
		flag = 0
	}

	game.isFullscreen = !game.isFullscreen
	game.window.SetFullscreen(flag)
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

	// Setup the other systems
	game.sManager.AddSystem(systems.NewInputSystem(game))
	for _, system := range game.sManager.Systems() {
		switch system.(type) {
		case *systems.InputSystem:
		}
	}
}
