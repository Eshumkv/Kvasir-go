package scenes

import (
	"github.com/eshumkv/Kvasir-go/components"
	"github.com/eshumkv/Kvasir-go/ecs"
)

func TestScene(mngr ecs.EntityManager) ecs.EntityManager {
	player := ecs.NewEntity(
		0, 0, 50, 50).Add(
		components.NewColorComponent(50, 60, 200)).Add(
		components.NewCollisionComponent())
	player.SetZ(11)
	test := ecs.NewEntity(
		100, -80, 50, 50).Add(
		components.NewColorComponent(5, 160, 100)).Add(
		components.NewCollisionComponent())

	mngr.Add(
		*player,
		"RenderSystem",
		"CameraSystem",
		"PlayerHandlingSystem",
		"CollisionSystem")
	mngr.Add(
		*test,
		"RenderSystem",
		"CollisionSystem")

	return mngr
}

func TestScene2(mngr ecs.EntityManager) ecs.EntityManager {
	player := ecs.NewEntity(
		0, 0, 50, 50).Add(
		components.NewColorComponent(100, 100, 80)).Add(
		components.NewCollisionComponent())
	player.SetZ(11)
	test := ecs.NewEntity(
		100, -80, 50, 50).Add(
		components.NewColorComponent(5, 160, 100)).Add(
		components.NewCollisionComponent())

	mngr.Add(
		*player,
		"RenderSystem",
		"CameraSystem",
		"PlayerHandlingSystem",
		"CollisionSystem")
	mngr.Add(
		*test,
		"RenderSystem",
		"CollisionSystem")

	return mngr
}
