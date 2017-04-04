package scenes

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/Eshumkv/kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

type MainMenuScene struct {
}

func NewMainMenuScene() *MainMenuScene {
	return &MainMenuScene{}
}

func (scene MainMenuScene) Init(world *ecs.World) {
	id := world.Create()
	world.AddComponents(id,
		components.NewRenderComponent(200, 100, 50),
		components.NewSpatialComponent(0, 0, 0, 50, 50),
		components.NewCameraFollowComponent())

	id = world.Create()
	world.AddComponents(id,
		components.NewRenderComponent(60, 100, 50),
		components.NewSpatialComponent(-50, -50, 0, 50, 50))
}

func (scene MainMenuScene) Dispose(world *ecs.World) {

}

func (scene MainMenuScene) Resume(world *ecs.World) {

}

func (scene MainMenuScene) Pause(world *ecs.World) {

}

func (scene MainMenuScene) Update(world *ecs.World, dt float64) {

	tempSystem := world.GetSystem("InputSystem")
	inputSystem := tempSystem.(*systems.InputSystem)

	tempSystem = world.GetSystem("SceneSystem")
	sceneSystem := tempSystem.(*systems.SceneSystem)

	if inputSystem.IsKeyDown(systems.CommandShoot) {
		sceneSystem.ChangeScene(NewTestScene2())
	}
}

func (scene MainMenuScene) Render(renderer *sdl.Renderer, world *ecs.World) {

}
