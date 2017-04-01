package scenes

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/Eshumkv/kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

type TestScene1 struct {
}

func NewTestScene1() *TestScene1 {
	return &TestScene1{}
}

func (scene TestScene1) Init(world *ecs.World) {
	id := world.Create()
	world.AddComponents(id,
		components.NewRenderComponent(200, 100, 50),
		components.NewSpatialComponent(0, 0, 0, 50, 50),
		components.NewCameraFollowComponent())
}

func (scene TestScene1) Dispose(world *ecs.World) {

}

func (scene TestScene1) Resume(world *ecs.World) {

}

func (scene TestScene1) Pause(world *ecs.World) {

}

func (scene TestScene1) Update(world *ecs.World, dt float64) {

	tempSystem := world.GetSystem("InputSystem")
	inputSystem := tempSystem.(*systems.InputSystem)

	tempSystem = world.GetSystem("SceneSystem")
	sceneSystem := tempSystem.(*systems.SceneSystem)

	if inputSystem.IsKeyDown(systems.CommandShoot) {
		sceneSystem.ChangeScene(NewTestScene2())
	}
}

func (scene TestScene1) Render(renderer *sdl.Renderer, world *ecs.World) {

}
