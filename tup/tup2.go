package tup

type Tup2[T any, U any] struct {
	first  T
	second U
}

func Of2[T, U any](first T, second U) Tup2[T, U] {
	return Tup2[T, U]{
		first:  first,
		second: second,
	}
}

func (p Tup2[T, U]) First() T {
	return p.first
}

func (p Tup2[T, U]) Second() U {
	return p.second
}

func (p Tup2[T, U]) Get() (T, U) {
	return p.first, p.second
}
