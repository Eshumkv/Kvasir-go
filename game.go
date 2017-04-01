package kvasir

import (
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/eshumkv/Kvasir-go/components"
	"github.com/eshumkv/Kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

// -----------------------------------------------------------------------------
// Game stuff

// Game struct that holds all relevant data
type Game struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	isRunning    bool
	isFullscreen bool
	world        ecs.World
}

// NewGame creates a new game.
func NewGame(window *sdl.Window, renderer *sdl.Renderer) *Game {
	result := &Game{
		window:    window,
		renderer:  renderer,
		isRunning: true,
	}

	// Init
	result.world = ecs.NewWorld([]ecs.SystemInterface{
		systems.NewInputSystem(result.Quit, result.ToggleFullscreen),
		systems.NewRenderSystem(result.renderer),
		systems.NewCameraSystem(result.window)})

	id := result.world.Create()
	result.world.AddComponents(id,
		components.NewRenderComponent(200, 100, 50),
		components.NewSpatialComponent(0, 0, 0, 50, 50),
		components.NewCameraFollowComponent())

	return result
}

// Update updates the gamestate.
func (game *Game) Update(dt float64) {
	game.world.SetDeltaTime(dt)
	game.world.Update()
}

// IsRunning checks whether the game is running.
func (game *Game) IsRunning() bool {
	return game.isRunning
}

// Quit allows you to quit the game.
func (game *Game) Quit() {
	game.isRunning = false
}

// ToggleFullscreen allows you to toggle whether the game displays
// fullscreen or not.
func (game *Game) ToggleFullscreen() {
	var flag uint32 = sdl.WINDOW_FULLSCREEN_DESKTOP

	if game.isFullscreen {
		flag = 0
	}

	game.isFullscreen = !game.isFullscreen
	game.window.SetFullscreen(flag)
}
