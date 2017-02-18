package systems

import (
	"github.com/eshumkv/Kvasir-go/ecs"
	"github.com/eshumkv/Kvasir-go/parts"
)

// CameraSystem defines the system to process input.
type CameraSystem struct {
	entity *ecs.Entity
	camera parts.CameraInterface
	mngr   *ecs.SystemManager
}

// NewCameraSystem returns a pointer to a new CameraSystem.
func NewCameraSystem(camera parts.CameraInterface) *CameraSystem {
	return &CameraSystem{
		camera: camera,
	}
}

// Init initializes the system.
func (s *CameraSystem) Init(mngr *ecs.SystemManager) {
	s.mngr = mngr
}

// Add adds an entity to the system.
func (s *CameraSystem) Add(e *ecs.Entity) {
	s.entity = e
	s.mngr.SendMessage(MessageCameraUpdate, s.camera)
}

// Update handles the update of the system.
func (s *CameraSystem) Update(dt float64) {
	if s.entity != nil {
		x, y := int32(s.entity.X()), int32(s.entity.Y())
		s.camera.SetX(x)
		s.camera.SetY(y)
	}
}

// Delete deletes an entity from this system.
func (s *CameraSystem) Delete(e ecs.Entity) {
	s.entity = nil
}

// HandleMessage handles any messages that need to be dealt with.
func (s CameraSystem) HandleMessage(msg ecs.Message, data interface{}) {
}
