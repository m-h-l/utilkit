package tup

type Tup3[T, U, V any] struct {
	first  T
	second U
	third  V
}

func Of3[T, U, V any](first T, second U, third V) Tup3[T, U, V] {
	return Tup3[T, U]{
		first:  first,
		second: second,
		third:  third,
	}
}

func (p Tup3[T, U, V]) First() T {
	return p.first
}

func (p Tup3[T, U, V]) Second() U {
	return p.second
}

func (p Tup3[T, U, V]) Third() U {
	return p.third
}

func (p Tup3[T, U, V]) Get() (T, U, V) {
	return p.first, p.second, p.third
}
