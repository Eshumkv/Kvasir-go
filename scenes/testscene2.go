package scenes

import (
	"github.com/Eshumkv/kvasir-go/components"
	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type TestScene2 struct {
}

func NewTestScene2() *TestScene2 {
	return &TestScene2{}
}

func (scene TestScene2) Init(world *ecs.World) {
	world.ClearEntities()

	id := world.Create()
	world.AddComponents(id,
		components.NewRenderComponent(0, 0, 255),
		components.NewSpatialComponent(70, 70, 0, 50, 50),
		components.NewCameraFollowComponent(),
		components.NewPlayerComponent())
	id = world.Create()
	world.AddComponents(id,
		components.NewRenderComponent(60, 100, 50),
		components.NewSpatialComponent(-50, -50, 0, 50, 50))
}

func (scene TestScene2) Dispose(world *ecs.World) {

}

func (scene TestScene2) Resume(world *ecs.World) {

}

func (scene TestScene2) Pause(world *ecs.World) {

}

func (scene TestScene2) Update(world *ecs.World, dt float64) {
}

func (scene TestScene2) Render(renderer *sdl.Renderer, world *ecs.World) {

}
