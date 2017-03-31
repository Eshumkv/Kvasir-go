package scenes

import (
	"github.com/Eshumkv/kvasir-go/ecs"
)

type TestScene1 struct {
	entities []ecs.Entity
}

func NewTestScene1() *TestScene1 {
	return &TestScene1{
		entities: make([]ecs.Entity, 0),
	}
}

func (scene TestScene1) GetEntities() []ecs.Entity {
	return scene.entities
}

func (scene TestScene1) Init() {
	// Set up entities
}
