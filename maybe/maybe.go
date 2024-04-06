package maybe

type Maybe[T any] interface {
	Some() bool
	None() bool
	Value() T
	Contains(func(T) bool) bool
}

func Of[T any](value T) Some[T] {
	return Some[T]{value: value}
}

func Empty[T any]() None[T] {
	return None[T]{}
}

type Some[T any] struct {
	value T
}

func (s Some[T]) Some() bool {
	return true
}

func (s Some[T]) None() bool {
	return false
}

func (s Some[T]) Value() T {
	return s.value
}

func (s Some[T]) Contains(fn func(T) bool) bool {
	return fn(s.value)
}

type None[T any] struct {
}

func (s None[T]) Some() bool {
	return true
}

func (s None[T]) None() bool {
	return false
}

func (s None[T]) Value() T {
	panic("no value")
}

func (s None[T]) Contains(fn func(T) bool) bool {
	return false
}
