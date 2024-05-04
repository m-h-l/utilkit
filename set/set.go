package set

type Set[T comparable] struct {
	a map[T]bool
}

func Of[T comparable](list []T) Set[T] {
	m := map[T]bool{}
	for _, item := range list {
		m[item] = true
	}
	return Set[T]{a: m}
}

func OfElements[T comparable](e ...T) Set[T] {
	return Of(e)
}

func Empty[T comparable]() Set[T] {
	return Set[T]{a: map[T]bool{}}
}

func (l Set[T]) Length() int {
	return len(l.a)
}

func (l Set[T]) Union(other Set[T]) Set[T] {
	m := map[T]bool{}
	for item := range l.a {
		m[item] = true
	}
	for item := range other.a {
		m[item] = true
	}
	return Set[T]{a: m}
}

func (l Set[T]) Add(item ...T) Set[T] {
	return l.Union(Of(item))
}

func (l Set[T]) Filter(fn func(T) bool) Set[T] {
	r := Empty[T]()
	for item := range l.a {
		if fn(item) {
			r.Add(item)
		}
	}
	return r
}

func (l Set[T]) ToSlice() []T {
	r := []T{}
	for item := range l.a {
		r = append(r, item)
	}
	return r
}

func Map[T comparable, U comparable](list Set[T], fn func(T) U) Set[U] {
	r := []U{}
	for item := range list.a {
		r = append(r, fn(item))
	}
	return Of(r)
}

func Fold[T comparable, U comparable](set Set[T], neutral U, fn func(U, T) U) U {
	slice := set.ToSlice()
	switch set.Length() {
	case 0:
		return neutral
	case 1:
		return fn(neutral, slice[0])
	}
	return Fold(Of(slice[1:]), fn(neutral, slice[0]), fn)
}
