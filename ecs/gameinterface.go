package ecs

// GameInterface defines the interface that every game should follow.
// Mostly used to decouple stuff so I can cyclic import :)
type GameInterface interface {
	Update(dt float64)
	Render(lag float64)
	IsRunning() bool
	SetRunning(state bool)
	ToggleFullscreen()
	HandleMessage(msg Message, data interface{}) interface{}
}
