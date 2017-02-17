package ecs

import "sort"

// SystemManager manages all the systems.
type SystemManager struct {
	systems []System
}

// NewSystemManager makes a new system manager.
func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: make([]System, 0),
	}
}

// Systems returns the systems that this manager manages.
func (mngr *SystemManager) Systems() []System {
	return mngr.systems
}

// AddSystem adds a system to the manager and sorts the list based on priority.
func (mngr *SystemManager) AddSystem(s System) {
	mngr.systems = append(mngr.systems, s)
	sort.Sort(byPriority(mngr.systems))
}

// Update goes through the sorted list and updates all the systems.
func (mngr *SystemManager) Update(dt float64) {
	for _, system := range mngr.systems {
		system.Update(dt)
	}
}
