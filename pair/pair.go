package pair

type Pair[T any, U any] struct {
	first  T
	second U
}

func Of[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		first:  first,
		second: second,
	}
}

func (p Pair[T, U]) First() T {
	return p.first
}

func (p Pair[T, U]) Second() U {
	return p.second
}

func (p Pair[T, U]) Get() T, U {
	return p.first, p.second
}
