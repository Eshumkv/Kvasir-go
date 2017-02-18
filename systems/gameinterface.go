package systems

// GameInterface defines the interface that every game should follow.
// Mostly used to decouple stuff so I can cyclic import :)
type GameInterface interface {
	Update(dt float64)
	Render(lag float64)
	IsRunning() bool
	SetRunning(state bool)
	SetCommand(c Command, state bool)
	GetCommand(c Command) bool
	ToggleFullscreen()
}
