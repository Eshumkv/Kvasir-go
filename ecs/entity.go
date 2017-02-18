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
		w:          w,
		h:          h,
		components: make([]Component, 0),
	}
	result.setRect()
	return result
}

// ID returns the unique id of the entity.
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

func (e *Entity) X() float64 {
	return e.x
}

func (e *Entity) SetX(n float64) {
	e.x = n
	e.setRect()
}

func (e *Entity) Y() float64 {
	return e.x
}

func (e *Entity) SetY(n float64) {
	e.y = n
	e.setRect()
}

func (e *Entity) W() int32 {
	return e.w
}

func (e *Entity) SetW(n int32) {
	e.w = n
	e.setRect()
}

func (e *Entity) H() int32 {
	return e.h
}

func (e *Entity) SetH(n int32) {
	e.h = n
	e.setRect()
}

func (e *Entity) Rect() sdl.Rect {
	return e.rect
}

func (e *Entity) setRect() {
	e.rect = sdl.Rect{
		X: int32(e.x),
		Y: int32(e.y),
		W: e.w,
		H: e.h,
	}
}
