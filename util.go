package utilkit

type Maybe[T any] interface {
	Some() bool
	None() bool
	Value() T
	Contains(func(T) bool) bool
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
	panic("no value in none")
}

func (s None[T]) Contains(fn func(T) bool) bool {
	return false
}

func First[T any](list []T) Maybe[T] {
	if len(list) >= 1 {
		return Some[T]{list[0]}
	}
	return None[T]{}
}

func Second[T any](list []T) Maybe[T] {
	if len(list) >= 2 {
		return Some[T]{list[1]}
	}
	return None[T]{}
}

func Third[T any](list []T) Maybe[T] {
	if len(list) >= 3 {
		return Some[T]{list[2]}
	}
	return None[T]{}
}

func Fourth[T any](list []T) Maybe[T] {
	if len(list) >= 4 {
		return Some[T]{list[3]}
	}
	return None[T]{}
}

func FindLast[T any](list []T, fn func(T) bool) (bool, int) {
	ok := false
	pos := -1
	for n, v := range list {
		if fn(v) {
			ok = true
			pos = n
		}
	}
	return ok, pos
}
