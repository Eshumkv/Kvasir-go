package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
)

type PlayerHandlingSystem struct {
	id       string
	entities []ecs.Entity
	mngr     *ecs.SystemManager
}

// NewPlayerHandlingSystem returns a pointer to a new PlayerHandlingSystem.
func NewPlayerHandlingSystem() *PlayerHandlingSystem {
	return &PlayerHandlingSystem{
		id:       "PlayerHandlingSystem",
		entities: make([]ecs.Entity, 0),
	}
}

// Init initializes the system.
func (s *PlayerHandlingSystem) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Update handles the update of the system.
func (s *PlayerHandlingSystem) Update(dt float64) {
	ret, err := s.mngr.SendMessage(MessageGetCommands, nil)
	if err != nil {
		return

	}
	commands := ret.([]bool)

	ret, err = s.mngr.SendMessage(MessageGetEntitiesOfSystem, s.id)
	if err != nil {
		return
	}
	entities := ret.([]ecs.Entity)

	speed := 100.0
	for _, entity := range entities {
		xAdd, yAdd := 0.0, 0.0
		if commands[CommandLeft] {
			xAdd = -speed * dt
		}
		if commands[CommandRight] {
			xAdd = speed * dt
		}

		if commands[CommandUp] {
			yAdd = -speed * dt
		}
		if commands[CommandDown] {
			yAdd = speed * dt
		}

		if xAdd != 0.0 || yAdd != 0.0 {
			newValueX := entity.X() + xAdd
			newValueY := entity.Y() + yAdd
			s.mngr.SendMessage(
				MessageSetEntityLocation,
				struct {
					id     uint64
					system string
					x      float64
					y      float64
				}{entity.ID(), s.id, newValueX, newValueY})
		}
	}
}

// Priority defines the priority of this system.
func (s PlayerHandlingSystem) Priority() uint {
	return 50
}

// HandleMessage handles any messages that need to be dealt with.
func (s *PlayerHandlingSystem) HandleMessage(
	msg ecs.Message, data interface{}) interface{} {
	return nil
}
