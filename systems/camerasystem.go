package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/eshumkv/Kvasir-go/parts"
)

// CameraSystem defines the system to process input.
type CameraSystem struct {
	id     string
	camera parts.CameraInterface
	mngr   *ecs.SystemManager
}

// NewCameraSystem returns a pointer to a new CameraSystem.
func NewCameraSystem(camera parts.CameraInterface) *CameraSystem {
	return &CameraSystem{
		id:     "CameraSystem",
		camera: camera,
	}
}

// Init initializes the system.
func (s *CameraSystem) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Add adds an entity to the system.
func (s *CameraSystem) Add(e *ecs.Entity) {
	panic("Not allowed to add to this system!")
}

// Update handles the update of the system.
func (s *CameraSystem) Update(dt float64) {

	ret, err := s.mngr.SendMessage(MessageGetEntitiesOfSystem, s.id)
	if err != nil {
		return
	}
	entities := ret.([]ecs.Entity)
	if len(entities) == 0 {
		return
	}
	entity := entities[0]
	x, y := int32(entity.X()), int32(entity.Y())

	s.camera.SetX(x)
	s.camera.SetY(y)
}

// Delete deletes an entity from this system.
func (s *CameraSystem) Delete(e ecs.Entity) {
	panic("Not allowed to delete from this system!")
}

// HandleMessage handles any messages that need to be dealt with.
func (s CameraSystem) HandleMessage(
	msg ecs.Message, data interface{}) interface{} {
	return nil
}
