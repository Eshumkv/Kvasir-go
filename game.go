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
	isFullscreen bool
	systems      *ecs.SystemManager
}

// NewGame creates a new game.
func NewGame(window *sdl.Window, renderer *sdl.Renderer) *Game {
	result := Game{
		window:    window,
		renderer:  renderer,
		isRunning: true,
		systems:   ecs.NewSystemManager(),
	}

	// Init
	renderer.SetDrawColor(110, 132, 174, 255)

	result.setupSystems()

	return &result
}

// BeforeUpdate runs before the update-loop. Do things you only want to do once
// per game-loop/frame/...
func (game *Game) BeforeUpdate(dt float64) {
	game.systems.BeforeUpdate(dt)
}

// Update updates the gamestate.
func (game *Game) Update(dt float64) {
	game.systems.Update(dt)
}

// Render shows the gamestate on screen.
func (game *Game) Render(lag float64) {
	game.renderer.Clear()
	game.systems.Render(lag)
	game.renderer.Present()
}

func (game *Game) IsRunning() bool {
	return game.isRunning
}

func (game *Game) SetRunning(state bool) {
	game.isRunning = state
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
	// Setup the systems
	game.systems.AddSystem(
		systems.NewRenderSystem(game.renderer), ecs.STypeRender)
	game.systems.AddSystem(systems.NewInputSystem(game), ecs.STypeBeforeUpdate)

	for _, system := range game.systems.AllSystems() {
		switch system.(type) {
		case *systems.RenderSystem:
			system.Init(game.systems)
			entity := ecs.NewEntity(
				0, 0, 50, 50).Add(
				components.NewColorComponent(50, 60, 200))
			system.Add(entity)
		case *systems.InputSystem:
			system.Init(game.systems)
		}
	}
}
