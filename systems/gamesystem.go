package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
)

// GameSystem defines the system to access the game objects.
type GameSystem struct {
	id   string
	mngr *ecs.SystemManager
	game *ecs.GameInterface
}

// NewGameSystem returns a pointer to a new GameSystem.
func NewGameSystem(game *ecs.GameInterface) *GameSystem {
	return &GameSystem{
		id:   "GameSystem",
		game: game,
	}
}

// Init initializes the system.
func (s *GameSystem) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Update handles the update of the system.
func (s *GameSystem) Update(dt float64) {
	// Empty
}

// Priority defines the priority of this system.
func (s GameSystem) Priority() uint {
	return 500
}

// HandleMessage handles any messages that need to be dealt with.
func (s *GameSystem) HandleMessage(
	msg ecs.Message, data interface{}) interface{} {

	switch msg {
	case MessageGetGameSnapshot:
		return s.game
	}
	return nil
}
