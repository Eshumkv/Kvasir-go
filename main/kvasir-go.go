package main

import (
	"time"

	kvasir "github.com/eshumkv/Kvasir-go"
	"github.com/veandco/go-sdl2/sdl"
)

// MsPerUpdate specifies the amount of milliseconds per update cycle is the ideal.
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
	lag := 0.0
	game := kvasir.NewGame(window, renderer)

	for game.IsRunning() {
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current
		lag += elapsed.Seconds()

		game.BeforeUpdate(elapsed.Seconds())

		for i := 0; i >= MaxNumUpdates && lag >= MsPerUpdate; i++ {
			game.Update(elapsed.Seconds())
			lag -= MsPerUpdate
		}

		game.Render(lag / MsPerUpdate)
	}
}
