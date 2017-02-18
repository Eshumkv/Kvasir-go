package systems

// Command defines the type used for the Command enum
type Command int

// The Command enum
const (
	CommandNone Command = iota
	CommandFullscreen
	CommandShoot
	CommandLeft
	CommandUp
	CommandRight
	CommandDown
	CommandSpeedup
	CommandCount
)
