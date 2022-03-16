package genericorderedmap

import (
	"container/list"
)

type mapElement[T comparable, U any] struct {
	key   T
	value U
}

func newMapElement[T comparable, U any](key T, value U) *mapElement[T, U] {
	return &mapElement[T, U]{
		key:   key,
		value: value,
	}
}

type Map[T comparable, U any] struct {
	mp map[T]*list.Element
	l  *list.List
}

func NewMap[T comparable, U any]() *Map[T, U] {
	return &Map[T, U]{
		mp: make(map[T]*list.Element),
		l:  list.New(),
	}
}

func (m *Map[T, U]) Set(key T, value U) {
	ele, ok := m.mp[key]
	if ok {
		// Update the value
		ele.Value.(*mapElement[T, U]).value = value
	} else {
		// Create a new value
		newEle := newMapElement(key, value)
		e := m.l.PushBack(newEle)
		m.mp[key] = e
	}
}

// Get returns the key corresponding value if exists.
// Otherwise, the second return value will be false.
func (m *Map[T, U]) Get(key T) (U, bool) {
	ele, ok := m.mp[key]
	if !ok {
		var empty U
		return empty, false
	}
	value := ele.Value.(*mapElement[T, U]).value

	return value, true
}

// Delete returns true if it successfully delete the key corresponding value.
// If the value does not exist, it returns false.
func (m *Map[T, U]) Delete(key T) bool {
	ele, ok := m.mp[key]
	if ok {
		delete(m.mp, key)
		m.l.Remove(ele)
	}

	return ok
}

func (m *Map[T, U]) Len() int {
	return len(m.mp)
}