package scenes

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/Eshumkv/kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

type TestScene2 struct {
}

func NewTestScene2() *TestScene2 {
	return &TestScene2{}
}

func (scene TestScene2) Init(world *ecs.World) {
	id := world.Create()
	world.AddComponents(id,
		components.NewRenderComponent(0, 0, 50),
		components.NewSpatialComponent(0, 0, 0, 50, 50),
		components.NewCameraFollowComponent())
}

func (scene TestScene2) Dispose(world *ecs.World) {

}

func (scene TestScene2) Resume(world *ecs.World) {

}

func (scene TestScene2) Pause(world *ecs.World) {

}

func (scene TestScene2) Update(world *ecs.World, dt float64) {

	tempSystem := world.GetSystem("InputSystem")
	inputSystem := tempSystem.(*systems.InputSystem)

	if inputSystem.IsKeyDown(systems.CommandShoot) {
		println("Hello")
	}
}

func (scene TestScene2) Render(renderer *sdl.Renderer, world *ecs.World) {

}
