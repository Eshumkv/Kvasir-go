package kvasir

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/Eshumkv/kvasir-go/parts"
	"github.com/Eshumkv/kvasir-go/systems"
	"github.com/eshumkv/Kvasir-go/scenes"
	"github.com/veandco/go-sdl2/sdl"
)

// -----------------------------------------------------------------------------
// Game stuff

// Game struct that holds all relevant data
type Game struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	isRunning       bool
	isFullscreen    bool
	camera          parts.CameraInterface
	systems         []ecs.SystemInterface
	entities        map[uint64]*ecs.Entity
	entitySystemMap map[string][]uint64
	inputSystem     *systems.InputSystem
	currentScene    ecs.SceneInterface
}

// NewGame creates a new game.
func NewGame(window *sdl.Window, renderer *sdl.Renderer) *Game {
	w, h := window.GetSize()
	result := Game{
		window:          window,
		renderer:        renderer,
		isRunning:       true,
		camera:          newCamera(w, h),
		systems:         make([]ecs.SystemInterface, 0),
		entities:        make(map[uint64]*ecs.Entity),
		entitySystemMap: make(map[string][]uint64),
		currentScene:    scenes.NewTestScene1(),
	}

	// Init
	renderer.SetDrawColor(110, 132, 174, 255)

	result.setupSystems()

	// TODO: set up scenes
	result.AddEntity(
		components.NewSpatialComponent(),
		components.NewRenderComponent())

	return &result
}

// Update updates the gamestate.
func (game *Game) Update(dt float64) {
	for _, system := range game.systems {
		system.Update(nil, dt)
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

// AddEntity adds an entity to the game.
func (game *Game) AddEntity(components ...ecs.ComponentInterface) uint64 {
	entity := ecs.NewEntity(
		len(components),
		game.entityAddComponentEvent,
		game.entityRemoveComponentEvent)

	for _, component := range components {
		entity.Add(component)
	}

	game.entities[entity.ID()] = entity

	for _, system := range game.systems {
		if entity.Has(system.GetComponentNames()) {
			smap := game.entitySystemMap[system.GetSystemName()]
			game.entitySystemMap[system.GetSystemName()] = append(
				smap, entity.ID())
		}
	}

	return entity.ID()
}

func (game *Game) ChangeScene(newScene ecs.SceneInterface,
	entitiesToKeep ...ecs.Entity) {
	count := len(newScene.GetEntities()) + len(entitiesToKeep)
	game.entities = make(map[uint64]*ecs.Entity, count)

	for _, entity := range newScene.GetEntities() {
		game.entities[entity.ID()] = &entity
	}
	for _, entity := range entitiesToKeep {
		game.entities[entity.ID()] = &entity
	}
}

// GetKeyState polls the inputsystem for the keystate.
func (game Game) GetKeyState(command systems.Command) bool {
	return game.inputSystem.IsKeyDown(command)
}

func (game *Game) entityAddComponentEvent(entityID uint64,
	componentName string) {
	for _, system := range game.systems {
		if contains(system.GetComponentNames(), componentName) {
			a := game.entitySystemMap[system.GetSystemName()]

			for _, e := range game.entities {
				if e.ID() == entityID {
					return
				}
			}

			a = append(a, entityID)
		}
	}
}

func (game *Game) entityRemoveComponentEvent(entityID uint64,
	componentName string) {
	for _, system := range game.systems {
		if contains(system.GetComponentNames(), componentName) {
			a := game.entitySystemMap[system.GetSystemName()]

			index := -1
			for i := range a {
				if a[i] == entityID {
					index = i
					break
				}
			}

			if index == -1 {
				panic("Something went really wrong")
			}

			i := index
			a = append(a[:i], a[i+1:]...)
		}
	}
}

func (game *Game) setupSystems() {
	game.inputSystem = systems.NewInputSystem(game.Quit, game.ToggleFullscreen)

	systemsToAdd := [...]ecs.SystemInterface{
		game.inputSystem,
		systems.NewRenderSystem(game.renderer)}

	for _, system := range systemsToAdd {
		game.systems = append(game.systems, system)
		game.entitySystemMap[system.GetSystemName()] = make([]uint64, 0)
	}
}

type listOfStrings []string

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------------------
// Camera stuff

func newCamera(w, h int) *Camera {
	return &Camera{
		halfWidth:  int32(w / 2),
		halfHeight: int32(h / 2),
	}
}

// Camera defines the camera used in the game.
type Camera struct {
	x, y       int32
	halfWidth  int32
	halfHeight int32
}

// GetX gets the X coordinate of the camera.
func (camera Camera) GetX() int32 {
	return camera.x
}

// GetY gets the Y coordinate of the camera.
func (camera Camera) GetY() int32 {
	return camera.y
}

// SetX sets the X component of the camera.
func (camera *Camera) SetX(x int32) {
	camera.x = x
}

// SetY sets the Y component of the camera.
func (camera *Camera) SetY(y int32) {
	camera.y = y
}

// GetScreenLocation calculates the correct position for the entity.
func (camera Camera) GetScreenLocation(x, y float64) (int, int) {
	screenX := (int32(x) - camera.x) + camera.halfWidth
	screenY := (int32(y) - camera.y) + camera.halfHeight
	return int(screenX), int(screenY)
}
