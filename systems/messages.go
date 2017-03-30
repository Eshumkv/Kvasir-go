package systems

import "github.com/eshumkv/Kvasir-go/ecs"

// The Message enum
const (
	// System Messages
	MessageGeneric ecs.Message = iota
	MessageCameraUpdate
	MessageGetCommands
	MessageGetEntitiesOfSystem
	MessageSetEntityLocation
	MessageGetGameSnapshot

	// Game Messages
	MessageGameQuit
	MessageGameChangeScene

	MessageCount
)
