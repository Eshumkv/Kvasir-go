package systems

import "github.com/eshumkv/Kvasir-go/ecs"

// The Message enum
const (
	MessageGeneric ecs.Message = iota
	MessageCameraUpdate
	MessageGetCommands
	MessageGetEntitiesOfSystem
	MessageSetEntityLocation
	MessageCount
)
