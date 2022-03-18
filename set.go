package genericorderedmap

type Set[T comparable] struct {
	mp *Map[T, struct{}]
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		mp: NewMap[T, struct{}](),
	}
}

func (s *Set[T]) Set(key T) {
	s.mp.Set(key, struct{}{})
}

// Get returns the key corresponding value if exists.
// Otherwise, the second return value will be false.
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.mp.Get(value)

	return ok
}

// Delete returns true if it successfully delete the key corresponding value.
// If the value does not exist, it returns false.
func (s *Set[T]) Delete(value T) bool {
	return s.mp.Delete(value)
}

// Len returns the number of elements stores in the map.
func (s *Set[T]) Len() int {
	return s.mp.Len()
}

// Keys returns the slice of the keys.
func (s *Set[T]) Values() []T {
	values := make([]T, s.Len())

	ele := s.mp.l.Front()
	for i := 0; ele != nil; i++ {
		values[i] = ele.Value.(*mapElement[T, struct{}]).key
		ele = ele.Next()
	}

	return values
}

func (s *Set[T]) FromValues(values []T) {
	for _, value := range values {
		s.Set(value)
	}
}

// Merge merges the other maps to this map
func (s *Set[T]) Merge(sets ...Set[T]) {
	for _, st := range sets {
		ele := st.mp.l.Front()
		for i := 0; ele != nil; i++ {
			value := ele.Value.(*mapElement[T, struct{}]).key
			s.Set(value)
		}
	}

}

func (s *Set[T]) Front() *SetElement[T] {
	front := s.mp.l.Front()

	return newSetElement[T](front)
}

func (s *Set[T]) Back() *SetElement[T] {
	back := s.mp.l.Back()

	return newSetElement[T](back)
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	st := NewSet[T]()
	st.Merge(*s, *other)

	return st
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	st := NewSet[T]()
	if s.Len() <= other.Len() {
		for _, value := range s.Values() {
			if other.Contains(value) {
				st.Set(value)
			}
		}
	} else {
		for _, value := range other.Values() {
			if s.Contains(value) {
				st.Set(value)
			}
		}
	}

	return st
}
