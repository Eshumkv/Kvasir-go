package main

import (
	"time"

	k "github.com/eshumkv/kvasir-go"
	"github.com/veandco/go-sdl2/sdl"
)

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
	lag := 0.0
	game := k.NewGame(window, renderer)

	for game.IsRunning {
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current
		lag += elapsed.Seconds()

		game.ProcessInput()
		for i := 0; i < k.MaxNumUpdates && lag >= k.MsPerUpdate; i++ {
			game.Update(elapsed.Seconds())
			lag -= k.MsPerUpdate
		}
		game.Render(lag / k.MsPerUpdate)
	}
}
