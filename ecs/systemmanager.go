package ecs

import (
	"errors"
	"sort"
)

// SystemType defines the type used for the SystemType enum
type SystemType int

// The SystemType enum
const (
	STypeBeforeUpdate SystemType = iota
	STypeUpdate
	STypeRender
	STypeCount
)

// SystemManager manages all the systems.
type SystemManager struct {
	beforeUpdateSystems []System
	updateSystems       []System
	renderSystems       []System
	allSystems          []System
	entityManager       EntityManager
	game                GameInterface
}

// NewSystemManager makes a new system manager.
func NewSystemManager(game GameInterface) *SystemManager {
	return &SystemManager{
		beforeUpdateSystems: make([]System, 0),
		updateSystems:       make([]System, 0),
		renderSystems:       make([]System, 0),
		allSystems:          make([]System, 0),
		game:                game,
	}
}

// Systems returns the systems that this manager manages.
func (mngr *SystemManager) Systems(t SystemType) []System {
	switch t {
	case STypeBeforeUpdate:
		return mngr.beforeUpdateSystems
	case STypeUpdate:
		return mngr.updateSystems
	case STypeRender:
		return mngr.renderSystems
	}
	return nil
}

// AllSystems returns ALL the systems that this manager manages.
func (mngr *SystemManager) AllSystems() []System {
	return mngr.allSystems
}

// AddSystem adds a system to the manager and sorts the list based on priority.
func (mngr *SystemManager) AddSystem(s System, t SystemType) {
	switch t {
	case STypeBeforeUpdate:
		mngr.beforeUpdateSystems = append(mngr.beforeUpdateSystems, s)
		sort.Sort(byPriority(mngr.beforeUpdateSystems))
	case STypeUpdate:
		mngr.updateSystems = append(mngr.updateSystems, s)
		sort.Sort(byPriority(mngr.updateSystems))
	case STypeRender:
		mngr.renderSystems = append(mngr.renderSystems, s)
		sort.Sort(byPriority(mngr.renderSystems))
	}

	count := len(mngr.beforeUpdateSystems) +
		len(mngr.updateSystems) + len(mngr.renderSystems)
	mngr.allSystems = make([]System, 0, count)

	for _, system := range mngr.beforeUpdateSystems {
		mngr.allSystems = append(mngr.allSystems, system)
	}

	for _, system := range mngr.updateSystems {
		mngr.allSystems = append(mngr.allSystems, system)
	}

	for _, system := range mngr.renderSystems {
		mngr.allSystems = append(mngr.allSystems, system)
	}
}

// BeforeUpdate goes through the sorted list and updates all the systems.
func (mngr *SystemManager) BeforeUpdate(dt float64) {
	for _, system := range mngr.beforeUpdateSystems {
		system.Update(dt)
	}
}

// Update goes through the sorted list and updates all the systems.
func (mngr *SystemManager) Update(dt float64) {
	for _, system := range mngr.updateSystems {
		system.Update(dt)
	}
}

// Render goes through the sorted list and renders all the systems.
func (mngr *SystemManager) Render(lag float64) {
	for _, system := range mngr.renderSystems {
		system.Update(lag)
	}
}

// SendMessage sends a message to all systems, entityManager and Game.
func (mngr *SystemManager) SendMessage(
	msg Message, data interface{}) (interface{}, error) {
	for _, system := range mngr.AllSystems() {
		ret := system.HandleMessage(msg, data)
		if ret != nil {
			return ret, nil
		}
	}

	ret := mngr.entityManager.HandleMessage(msg, data)
	if ret != nil {
		return ret, nil
	}

	ret = mngr.game.HandleMessage(msg, data)
	if ret != nil {
		return ret, nil
	}

	return nil, errors.New("Whoopsie")
}

// SetEntityManager sets the entitymanager.
func (mngr *SystemManager) SetEntityManager(em EntityManager) {
	mngr.entityManager = em
}

// GetEntityManager returns the entitymanager.
func (mngr SystemManager) GetEntityManager() *EntityManager {
	return &mngr.entityManager
}
