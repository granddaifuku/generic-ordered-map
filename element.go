package genericorderedmap

import "container/list"

type Element[T comparable, U any] struct {
	element *list.Element
	Key     T
	Value   U
}

func newElement[T comparable, U any](e *list.Element) *Element[T, U] {
	if e == nil {
		return nil
	}

	ele := e.Value.(*mapElement[T, U])

	return &Element[T, U]{
		element: e,
		Key:     ele.key,
		Value:   ele.value,
	}
}

func (e *Element[T, U]) Next() *Element[T, U] {
	return newElement[T, U](e.element.Next())
}

func (e *Element[T, U]) Prev() *Element[T, U] {
	return newElement[T, U](e.element.Prev())
}

type SetElement[T comparable] struct {
	element *list.Element
	Value   T
}

func newSetElement[T comparable](e *list.Element) *SetElement[T] {
	if e == nil {
		return nil
	}

	ele := e.Value.(*mapElement[T, struct{}])

	return &SetElement[T]{
		element: e,
		Value:   ele.key,
	}
}

func (se *SetElement[T]) Next() *SetElement[T] {
	return newSetElement[T](se.element.Next())
}

func (se *SetElement[T]) Prev() *SetElement[T] {
	return newSetElement[T](se.element.Prev())
}
