package list

import (
	"github.com/m-h-l/utilkit/maybe"
)

type List[T any] struct {
	a []T
}

func Of[T any](list []T) List[T] {
	return List[T]{a: list}
}

func OfElements[T any](e ...T) List[T] {
	return List[T]{a: e}
}

func Empty[T any]() List[T] {
	return List[T]{a: []T{}}
}

func (l List[T]) Length() int {
	return len(l.a)
}

func (l List[T]) Head() maybe.Maybe[T] {
	if len(l.a) >= 1 {
		return maybe.Of(l.a[0])
	}
	return maybe.Empty[T]()
}

func (l List[T]) Last() maybe.Maybe[T] {
	if len(l.a) >= 1 {
		return maybe.Of(l.a[len(l.a)-1])
	}
	return maybe.Empty[T]()
}

func (l List[T]) Tail() maybe.Maybe[List[T]] {
	if len(l.a) >= 2 {
		return maybe.Of(Of(l.a[1:]))
	}
	return maybe.Empty[List[T]]()
}

func (l List[T]) At(idx int) maybe.Maybe[T] {
	if len(l.a) >= idx {
		return maybe.Of(l.a[idx])
	}
	return maybe.Empty[T]()
}

func (l List[T]) Append(item ...T) List[T] {
	return Of(append(l.a, item...))
}

func (l List[T]) Filter(fn func(T) bool) List[T] {
	r := Empty[T]()
	for _, v := range l.a {
		if fn(v) {
			r.Append(v)
		}
	}
	return r
}

func (l List[T]) ToSlice() []T {
	return l.a
}

func Map[T any, U any](list List[T], fn func(T) U) List[U] {
	r := []U{}
	for _, item := range list.a {
		r = append(r, fn(item))
	}
	return Of(r)
}

func FoldLeft[T any, U any](list List[T], neutral U, fn func(U, T) U) U {
	switch list.Length() {
	case 0:
		return neutral
	case 1:
		return fn(neutral, list.a[0])
	}
	return FoldLeft(Of(list.a[1:]), fn(neutral, list.a[0]), fn)
}
