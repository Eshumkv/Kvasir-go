package ecs

type SceneInterface interface {
	GetEntities() []Entity
	Init()
}
