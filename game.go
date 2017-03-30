package kvasir

import (
	"github.com/Eshumkv/kvasir-go/parts"
	"github.com/eshumkv/Kvasir-go/ecs"
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
	camera       parts.CameraInterface
	systems      []ecs.SystemInterface
}

// NewGame creates a new game.
func NewGame(window *sdl.Window, renderer *sdl.Renderer) *Game {
	w, h := window.GetSize()
	result := Game{
		window:    window,
		renderer:  renderer,
		isRunning: true,
		camera:    newCamera(w, h),
		systems:   make([]ecs.SystemInterface, 0),
	}

	// Init
	renderer.SetDrawColor(110, 132, 174, 255)

	result.setupSystems()

	return &result
}

// Update updates the gamestate.
func (game *Game) Update(dt float64) {
	for _, system := range game.systems {
		system.Update(dt)
	}
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

func (game *Game) setupSystems() {
	systemsToAdd := [...]ecs.SystemInterface{
		systems.NewInputSystem(game.Quit),
		systems.NewRenderSystem(game.renderer)}

	for _, system := range systemsToAdd {
		game.systems = append(game.systems, system)
	}
}

// -----------------------------------------------------------------------------
// Camera stuff

func newCamera(w, h int) *Camera {
	return &Camera{
		halfWidth:  int32(w / 2),
		halfHeight: int32(h / 2),
	}
}

// Camera implements the parts.CameraInterface.
type Camera struct {
	x, y       int32
	halfWidth  int32
	halfHeight int32
}

func (camera Camera) GetX() int32 {
	return camera.x
}

func (camera Camera) GetY() int32 {
	return camera.y
}

func (camera *Camera) SetX(x int32) {
	camera.x = x
}
func (camera *Camera) SetY(y int32) {
	camera.y = y
}

func (camera Camera) GetScreenLocation(x, y float64) (int, int) {
	screenX := (int32(x) - camera.x) + camera.halfWidth
	screenY := (int32(y) - camera.y) + camera.halfHeight
	return int(screenX), int(screenY)
}
