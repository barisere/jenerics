package jenerics

type Semigroup[T any] interface {
	Concat(T, T) T
}

type Monoid[T any] interface {
	Semigroup[T]
	Identity() T
}
