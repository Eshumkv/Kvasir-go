package main

import (
	"time"

	"github.com/Eshumkv/kvasir-go/ecs"
	"github.com/Eshumkv/kvasir-go/scenes"
	"github.com/Eshumkv/kvasir-go/systems"
	"github.com/veandco/go-sdl2/sdl"
)

// MsPerUpdate specifies the amount of milliseconds per update cycle
// is the ideal.
const MsPerUpdate = 1 / 100.0

// MaxNumUpdates specifies how many times update should be called to "catch up".
const MaxNumUpdates = 5

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Kvasir",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		1280, 720,
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	defer renderer.Destroy()

	previous := time.Now()

	isFullscreen := false
	isRunning := true

	quitFunc := func() {
		isRunning = false
	}
	fullscreenFunc := func() {
		var flag uint32 = sdl.WINDOW_FULLSCREEN_DESKTOP

		if isFullscreen {
			flag = 0
		}

		isFullscreen = !isFullscreen
		window.SetFullscreen(flag)
	}

	// Init
	world := ecs.NewWorld([]ecs.SystemInterface{
		systems.NewInputSystem(quitFunc, fullscreenFunc),
		systems.NewRenderSystem(renderer),
		systems.NewCameraSystem(window),
		systems.NewSceneSystem(renderer),
		systems.NewPlayerSystem()})
	tempSystem := world.GetSystem("SceneSystem")
	sceneSystem := tempSystem.(*systems.SceneSystem)
	sceneSystem.SetFirstScene(scenes.NewMainMenuScene(), &world)

	for isRunning {
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current

		world.SetDeltaTime(elapsed.Seconds())
		world.Update()
	}
}
