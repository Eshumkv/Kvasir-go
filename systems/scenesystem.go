package systems

import (
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------------------------------------------------
// SceneSystem

// SceneSystem is the system that handles the camera.
type SceneSystem struct {
	systemName   string
	renderer     *sdl.Renderer
	currentScene SceneInterface
	nextScene    SceneInterface
}

// NewSceneSystem creates a new SceneSystem
func NewSceneSystem(renderer *sdl.Renderer) *SceneSystem {
	return &SceneSystem{
		systemName: "SceneSystem",
		renderer:   renderer,
	}
}

// SetFirstScene sets the starting scene. Call this after creating the world.
func (system *SceneSystem) SetFirstScene(
	scene SceneInterface, world *ecs.World) {

	system.currentScene = scene
	system.currentScene.Init(world)
	system.nextScene = nil
}

// Update updates this system.
func (system *SceneSystem) Update(
	entities []ecs.EntityID, world *ecs.World, dt float64) {

	system.currentScene.Update(world, dt)

	if system.nextScene != nil {
		// TODO: Some transition here?
		system.currentScene = system.nextScene
		system.currentScene.Init(world)
		system.nextScene = nil
	}
}

// GetComponentNames gives a list of components that this system uses.
func (system SceneSystem) GetComponentNames() []string {
	var d []string
	return d
}

// GetSystemName returns the name of this system.
func (system SceneSystem) GetSystemName() string {
	return system.systemName
}

// ChangeScene changes the scene at the next convenient moment.
func (system *SceneSystem) ChangeScene(scene SceneInterface) {
	system.nextScene = scene
}

// SceneInterface defines the interface for a scene.
type SceneInterface interface {
	Init(world *ecs.World)
	Dispose(world *ecs.World)
	Resume(world *ecs.World)
	Pause(world *ecs.World)
	Update(world *ecs.World, dt float64)
	Render(renderer *sdl.Renderer, world *ecs.World)
}
