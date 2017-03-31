package ecs

// SceneManagerInterface defines the interface for a scene manager.
type SceneManagerInterface interface {
	AddEntity(components ...ComponentInterface) uint64
	ChangeScene(newScene SceneInterface, entitiesToKeep ...Entity)
}
