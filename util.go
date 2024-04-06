package utilkit

import (
	"github.com/m-h-l/utilkit/maybe"
)

func First[T any](list []T) maybe.Maybe[T] {
	if len(list) >= 1 {
		return maybe.Of[T](list[0])
	}
	return maybe.None[T]{}
}

func Second[T any](list []T) maybe.Maybe[T] {
	if len(list) >= 2 {
		return maybe.Of[T](list[1])
	}
	return maybe.Empty[T]()
}

func Third[T any](list []T) maybe.Maybe[T] {
	if len(list) >= 3 {
		return maybe.Of[T](list[2])
	}
	return maybe.Empty[T]()
}

func Fourth[T any](list []T) maybe.Maybe[T] {
	if len(list) >= 4 {
		return maybe.Of[T](list[3])
	}
	return maybe.Empty[T]()
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
