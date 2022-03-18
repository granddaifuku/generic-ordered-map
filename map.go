package genericorderedmap

import (
	"container/list"
)

type Entry[T comparable, U any] struct {
	Key   T
	Value U
}

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

// Len returns the number of elements stores in the map.
func (m *Map[T, U]) Len() int {
	return len(m.mp)
}

// Keys returns the slice of the keys.
func (m *Map[T, U]) Keys() []T {
	keys := make([]T, m.Len())

	ele := m.l.Front()
	for i := 0; ele != nil; i++ {
		keys[i] = ele.Value.(*mapElement[T, U]).key
		ele = ele.Next()
	}

	return keys
}

// Values returns the slice of the values.
func (m *Map[T, U]) Values() []U {
	values := make([]U, m.Len())

	ele := m.l.Front()
	for i := 0; ele != nil; i++ {
		values[i] = ele.Value.(*mapElement[T, U]).value
		ele = ele.Next()
	}

	return values
}

// Entries converts map to slice of Entry
func (m *Map[T, U]) Entries() []Entry[T, U] {
	entries := make([]Entry[T, U], m.Len())

	ele := m.l.Front()
	for i := 0; ele != nil; i++ {
		key := ele.Value.(*mapElement[T, U]).key
		value := ele.Value.(*mapElement[T, U]).value
		entries[i] = Entry[T, U]{Key: key, Value: value}
		ele = ele.Next()
	}

	return entries
}

func (m *Map[T, U]) FromEntries(entries []Entry[T, U]) {
	for _, entry := range entries {
		m.Set(entry.Key, entry.Value)
	}
}

// Merge merges the other maps to this map
func (m *Map[T, U]) Merge(maps ...Map[T, U]) {
	for _, mp := range maps {
		entry := mp.Entries()
		m.FromEntries(entry)
	}
}

func (m *Map[T, U]) Front() *Element[T, U] {
	front := m.l.Front()
	if front == nil {
		return nil
	}

	return newElement[T, U](front)
}

func (m *Map[T, U]) Back() *Element[T, U] {
	back := m.l.Back()
	if back == nil {
		return nil
	}

	return newElement[T, U](back)
}
