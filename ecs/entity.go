package ecs

import (
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	counterLock sync.Mutex
	idInc       uint64
)

// Entity defines the entity in the entity component system.
type Entity struct {
	id         uint64
	x, y       float64
	z          int32
	w, h       int32
	components []Component
	rect       sdl.Rect
}

// NewEntity creates a new entity with a unique identifier.
func NewEntity(x, y float64, w, h int32) *Entity {
	result := &Entity{
		id:         atomic.AddUint64(&idInc, 1),
		x:          x,
		y:          y,
		z:          10,
		w:          w,
		h:          h,
		components: make([]Component, 0),
	}
	return result
}

// ID returns the unique id of the entity. (identifier interface)
func (e Entity) ID() uint64 {
	return e.id
}

// Add a component to the entity.
func (e *Entity) Add(comp Component) *Entity {
	e.components = append(e.components, comp)
	return e
}

// Remove a component from the entity.
func (e *Entity) Remove(comp string) {
	index := -1
	for i, c := range e.components {
		if reflect.TypeOf(c).Name() == comp {
			index = i
			break
		}
	}

	if index != -1 {
		copy(e.components[index:], e.components[index+1:])
		e.components[len(e.components)-1] = nil
		e.components = e.components[:len(e.components)-1]
	}
}

// Get the component that you want, returns (component, ok)
func (e Entity) Get(comp string) (Component, bool) {
	for i := range e.components {
		if reflect.TypeOf(e.components[i]).String() == comp {
			return e.components[i], true
		}
	}
	return nil, false
}

// Has the entity a specific component?
func (e Entity) Has(comp string) bool {
	for _, c := range e.components {
		if reflect.TypeOf(c).Name() == comp {
			return true
		}
	}
	return false
}

// X returns the x value of the entity.
func (e Entity) X() float64 {
	return e.x
}

// SetX sets the x value of the entity.
func (e *Entity) SetX(n float64) {
	e.x = n
}

// Y returns the y value of the entity.
func (e Entity) Y() float64 {
	return e.y
}

// SetY sets the y value of the entity.
func (e *Entity) SetY(n float64) {
	e.y = n
}

// Z returns the z value of the entity.
func (e Entity) Z() int32 {
	return e.z
}

// SetZ sets the z value of the entity.
func (e *Entity) SetZ(n int32) {
	e.z = n
}

// W returns the width of the entity.
func (e Entity) W() int32 {
	return e.w
}

// SetW sets the width of the entity.
func (e *Entity) SetW(n int32) {
	e.w = n
}

// H returns the height of the entity.
func (e Entity) H() int32 {
	return e.h
}

// SetH sets the height of the entity.
func (e *Entity) SetH(n int32) {
	e.h = n
}

// Rect returns a (SDL) rectangle that defines this entity.
func (e *Entity) Rect() sdl.Rect {
	return sdl.Rect{
		X: int32(e.x),
		Y: int32(e.y),
		W: e.w,
		H: e.h,
	}
}
