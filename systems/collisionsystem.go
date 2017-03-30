package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
)

type CollisionSystem struct {
	id   string
	mngr *ecs.SystemManager
}

// NewPlayerHandlingSystem returns a pointer to a new PlayerHandlingSystem.
func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{
		id: "CollisionSystem",
	}
}

// Init initializes the system.
func (s *CollisionSystem) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Update handles the update of the system.
func (s *CollisionSystem) Update(dt float64) {
	// ret, err := s.mngr.SendMessage(MessageGetEntitiesOfSystem, s.id)
	// if err != nil {
	// 	return
	// }
	// entities := ret.([]ecs.Entity)
	// for _, entity := range entities {
	// 	fmt.Println(entity)
	// }
}

// Priority defines the priority of this system.
func (s CollisionSystem) Priority() uint {
	return 100
}

// HandleMessage handles any messages that need to be dealt with.
func (s *CollisionSystem) HandleMessage(
	msg ecs.Message, data interface{}) interface{} {
	return nil
}
